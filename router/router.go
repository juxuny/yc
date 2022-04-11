package router

import (
	"context"
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/utils"
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
		log.Debug("register method: ", path)
	}
	return nil
}

func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	start := time.Now()
	statusCode := http.StatusOK
	defer func(s time.Time) {
		log.Infof("%d %v %v [%v]", statusCode, r.Method, path, time.Now().Sub(s))
	}(start)
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
		WriteJson(w, errors.SystemError.NotFound, http.StatusNotFound)
		return
	}
	ctx := context.Background()
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		WriteJson(w, errors.SystemError.NotSupportedMethod, http.StatusBadRequest)
		return
	}
	in := []reflect.Value{reflect.ValueOf(ctx)}
	tt := caller.Type()
	ct := r.Header.Get("Content-Type")
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
			WriteJson(w, errors.SystemError.NotSupportedMethod, http.StatusBadRequest)
			return
		}
		_ = r.Body.Close()
		if err := json.Unmarshal(data, requestParamInstance); err != nil {
			WriteJson(w, errors.SystemError.InvalidJsonData.Wrap(err), http.StatusBadRequest)
			return
		}
	} else {
		if r.Method == http.MethodGet {
			if err := decoder.Decode(requestParamInstance, r.URL.Query()); err != nil {
				WriteJson(w, errors.SystemError.InvalidInputDataObject.Wrap(err), http.StatusBadRequest)
				return
			}
		} else {
			if err := r.ParseForm(); err != nil {
				WriteJson(w, errors.SystemError.InvalidFormData.Wrap(err), http.StatusBadRequest)
				return
			}
			if err := decoder.Decode(requestParamInstance, r.PostForm); err != nil {
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
			if err, ok := secondResponse.(error); ok {
				WriteJson(w, err, http.StatusBadRequest)
				return
			}
		}
		WriteSuccessData(w, responses[0].Interface())
	}
}