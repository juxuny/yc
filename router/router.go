package router

import (
	"context"
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/trace"
	"github.com/juxuny/yc/utils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

var decoder = schema.NewDecoder() //把get, post 请求转成 struct

type Router struct {
	Prefix         string
	m              map[string]reflect.Value
	RecoverHandler func(w http.ResponseWriter, r *http.Request)
}

func NewRouter(prefix string) *Router {
	return &Router{
		Prefix: prefix,
		m:      make(map[string]reflect.Value),
		RecoverHandler: func(w http.ResponseWriter, r *http.Request) {
		},
	}
}

func (t *Router) Register(groupName string, handler interface{}) error {
	vv := reflect.ValueOf(handler)
	tt := reflect.TypeOf(handler)
	for i := 0; i < vv.NumMethod(); i++ {
		method := vv.Method(i)
		methodType := tt.Method(i)
		if method.Type().NumIn() != 2 {
			return errors.SystemError.InvalidNumberOfParams
		}
		if method.Type().NumOut() != 2 {
			return errors.SystemError.InvalidNumberOfReplyEntities
		}
		methodName := strings.ReplaceAll(utils.ToUnderLine(methodType.Name), "_", "-")
		path := t.Prefix + "/" + groupName + "/" + methodName
		t.m[path] = method
		log.Info("register method: ", path)
	}
	return nil
}

func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	start := time.Now()
	statusCode := http.StatusOK
	ctx := context.Background()
	prevMetadata := CreateMetadataFromHeader(r.Header)
	ctx = metadata.NewIncomingContext(ctx, prevMetadata)
	trace.InitContext(prevMetadata)
	defer trace.Clean()
	defer func(s time.Time, httpStatus *int) {
		log.Infof("%d %v %v [%v]", *httpStatus, r.Method, path, time.Now().Sub(s))
	}(start, &statusCode)
	callFinished := make(chan bool, 1)
	trace.GoRun(func() {
		defer func() {
			close(callFinished)
		}()
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
				debug.PrintStack()
				t.RecoverHandler(w, r)
				return
			}
		}()
		caller, b := t.m[path]
		if !b {
			statusCode = http.StatusNotFound
			WriteJson(w, errors.SystemError.NotFound, http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPost && r.Method != http.MethodGet {
			statusCode = http.StatusBadRequest
			WriteJson(w, errors.SystemError.NotSupportedMethod, http.StatusBadRequest)
			return
		}
		md := trace.GetMetadata()
		ctx = metadata.NewIncomingContext(ctx, md)
		in := []reflect.Value{reflect.ValueOf(ctx)}
		tt := caller.Type()
		ct := r.Header.Get("Content-Type")
		if ct == "" {
			ct = r.Header.Get("content-type")
		}
		var requestParam reflect.Value
		if tt.In(1).Kind() == reflect.Ptr {
			requestParam = reflect.New(tt.In(1).Elem())
		} else {
			requestParam = reflect.New(tt.In(1))
		}
		requestParamInstance := requestParam.Interface()
		if r.Method == http.MethodPost && strings.Contains(ct, "json") {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				statusCode = http.StatusBadRequest
				WriteJson(w, errors.SystemError.NotSupportedMethod, http.StatusBadRequest)
				return
			}
			_ = r.Body.Close()
			log.Info("request:", strings.ReplaceAll(string(data), "\n", ""))
			if err := json.Unmarshal(data, requestParamInstance); err != nil {
				statusCode = http.StatusBadRequest
				WriteJson(w, errors.SystemError.InvalidJsonData.Wrap(err), http.StatusBadRequest)
				return
			}
		} else if r.Method == http.MethodPost && strings.Contains(ct, "protobuf") {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				statusCode = http.StatusBadRequest
				WriteJson(w, errors.SystemError.NotSupportedMethod, http.StatusBadRequest)
				return
			}
			_ = r.Body.Close()
			log.Info("request data size:", len(data))
			if message, ok := requestParamInstance.(proto.Message); ok {
				if err := proto.Unmarshal(data, message); err != nil {
					statusCode = http.StatusBadRequest
					WriteProtobufError(w, dt.FromError(errors.SystemError.InvalidProtobufData.Wrap(err)), http.StatusBadRequest)
					return
				}
			} else {
				statusCode = http.StatusBadRequest
				WriteProtobufError(w, dt.FromError(errors.SystemError.InvalidProtobufHolder.Wrap(err)), http.StatusBadRequest)
				return
			}
		} else {
			log.Debug("content-type:", ct)
			if r.Method == http.MethodGet {
				log.Info("request:", r.URL.Query().Encode())
				if err := decoder.Decode(requestParamInstance, r.URL.Query()); err != nil {
					statusCode = http.StatusBadRequest
					WriteJson(w, errors.SystemError.InvalidInputDataObject.Wrap(err), http.StatusBadRequest)
					return
				}
			} else {
				if err := r.ParseForm(); err != nil {
					statusCode = http.StatusBadRequest
					WriteJson(w, errors.SystemError.InvalidFormData.Wrap(err), http.StatusBadRequest)
					return
				}
				log.Info("request:", r.PostForm.Encode())
				if err := decoder.Decode(requestParamInstance, r.PostForm); err != nil {
					statusCode = http.StatusBadRequest
					WriteJson(w, errors.SystemError.InvalidInputDataObject.Wrap(err), http.StatusBadRequest)
					return
				}
			}
		}
		in = append(in, reflect.ValueOf(requestParamInstance))
		responses := caller.Call(in)
		if len(responses) == 0 {
			log.Error("no response")
		} else if len(responses) == 1 {
			WriteSuccessData(w, responses[0].Interface())
		} else if len(responses) == 2 {
			secondResponse := responses[1].Interface()
			if secondResponse != nil {
				statusCode = http.StatusBadRequest
				if err, ok := secondResponse.(errors.Error); ok {
					if IsProtobufContentType(ct) {
						WriteProtobufError(w, dt.FromError(err), statusCode)
					} else {
						WriteJson(w, err, statusCode)
					}
					return
				}
				if err, ok := secondResponse.(*errors.Error); ok {
					if IsProtobufContentType(ct) {
						WriteProtobufError(w, dt.FromError(*err), statusCode)
					} else {
						WriteJson(w, *err, statusCode)
					}
				}
				if err, ok := secondResponse.(error); ok {
					if IsProtobufContentType(ct) {
						WriteProtobufError(w, dt.FromError(errors.SystemError.InternalError.Wrap(err)), statusCode)
					} else {
						WriteJson(w, errors.SystemError.InternalError.Wrap(err), statusCode)
					}
					return
				}
			}
			if strings.Contains(ct, "protobuf") {
				if message, ok := responses[0].Interface().(proto.Message); ok {
					WriteProtobuf(w, message)
				} else {
					WriteSuccessData(w, responses[0].Interface())
				}
			} else {
				WriteSuccessData(w, responses[0].Interface())
			}
		}
	})
	<-callFinished
}
