package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/trace"
	"github.com/juxuny/yc/utils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

type HandleFunc func(ctx *Context)

func RecoverHandler(ctx *Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			debug.PrintStack()
			return
		}
	}()
	ctx.Next()
}

func TraceHandler(ctx *Context) {
	ctx.prevMetaData = CreateMetadataFromHeader(ctx.Request.Header)
	ctx.RpcContext = metadata.NewIncomingContext(context.Background(), ctx.prevMetaData)
	trace.InitContext(ctx.prevMetaData)
	defer trace.Clean()
	ctx.Next()
}

func LoggerHandler(ctx *Context) {
	start := time.Now()
	defer func() {
		log.Infof("%d %v %v [%v]", ctx.StatusCode, ctx.Request.Method, ctx.Request.URL.Path, time.Now().Sub(start))
	}()
}

type mainHandlerWrapper struct {
	m      map[string]reflect.Value
	prefix string
}

func newMainHandler(prefix string) *mainHandlerWrapper {
	return &mainHandlerWrapper{
		m:      make(map[string]reflect.Value),
		prefix: prefix,
	}
}

func (t *mainHandlerWrapper) initializeRequestInstance(ctx *Context, caller reflect.Value) (inputValue []reflect.Value, requestParamInstance interface{}) {
	inputValue = []reflect.Value{reflect.ValueOf(ctx.RpcContext)}
	tt := caller.Type()
	var requestParam reflect.Value
	if tt.In(1).Kind() == reflect.Ptr {
		requestParam = reflect.New(tt.In(1).Elem())
	} else {
		requestParam = reflect.New(tt.In(1))
	}
	requestParamInstance = requestParam.Interface()
	return
}

func (t *mainHandlerWrapper) process(ctx *Context, writeSuccess, writeFailed func(v interface{}) error) {
	if !ctx.IsPost() {
		return
	}
	path := ctx.Request.URL.Path
	caller, b := t.m[path]
	if !b {
		err := writeFailed(errors.SystemError.NotFound)
		if err != nil {
			log.Error(err)
			return
		}
	}

	in, requestParamInstance := t.initializeRequestInstance(ctx, caller)
	data, err := ctx.GetRequestBodyBytes()
	if err != nil {
		log.Error(err)
		err = writeFailed(errors.SystemError.HttpError.Wrap(err))
		if err != nil {
			log.Error(err)
			return
		}
		return
	}

	if err := json.Unmarshal(data, requestParamInstance); err != nil {
		err = writeFailed(errors.SystemError.InvalidJsonData.Wrap(err))
		if err != nil {
			log.Error(err)
		}
		return
	}
	in = append(in, reflect.ValueOf(requestParamInstance))
	responses := caller.Call(in)
	if len(responses) == 0 {
		log.Error("no response")
		return
	} else if len(responses) == 1 {
		err = writeSuccess(responses[0].Interface())
		if err != nil {
			log.Error(err)
			return
		}
	} else if len(responses) == 2 {
		secondResponse := responses[1].Interface()
		if secondResponse != nil {
			err = writeFailed(secondResponse)
			if err != nil {
				log.Error(err)
			}
		} else {
			err = writeSuccess(responses[0].Interface())
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (t *mainHandlerWrapper) callerHandleFunc(ctx *Context) {
	if ctx.IsJson() {
		t.process(ctx, func(v interface{}) error {
			_, err := ctx.WriteJsonSuccess(v)
			return err
		}, func(v interface{}) error {
			_, err := ctx.WriteJsonFailed(v)
			return err
		})
	} else if ctx.IsProtoBuf() {
		t.process(ctx, func(v interface{}) error {
			if msg, ok := v.(proto.Message); ok {
				_, err := ctx.WriteProtobufSuccess(msg)
				return err
			} else {
				_, err := ctx.WriteProtobufFailed(dt.FromError(errors.SystemError.InternalError.Wrap(fmt.Errorf("got an invalid response entitiy"))))
				return err
			}
		}, func(v interface{}) error {
			var err error
			if errorValue, ok := v.(errors.Error); ok {
				_, err = ctx.WriteProtobufFailed(dt.FromError(errorValue))
			} else if errorValue, ok := v.(*errors.Error); ok {
				_, err = ctx.WriteProtobufFailed(dt.FromError(*errorValue))
			} else if errorValue, ok := v.(error); ok {
				_, err = ctx.WriteProtobufFailed(dt.FromError(errors.SystemError.InternalError.Wrap(errorValue)))
			} else {
				_, err = ctx.WriteProtobufFailed(dt.FromError(errors.SystemError.InternalError.Wrap(fmt.Errorf("got an invalid response entity"))))
			}
			return err
		})
	}
}

func (t *mainHandlerWrapper) register(groupName string, handler interface{}) error {
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
		path := t.prefix + "/" + groupName + "/" + methodName
		t.m[path] = method
		log.Info("register method: ", path)
	}
	return nil
}
