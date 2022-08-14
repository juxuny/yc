package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	BodyContent    []byte

	prevMetaData metadata.MD
	RpcContext   context.Context
	handlerChain []HandleFunc
	index        int
	StatusCode   int
}

func newContext(handlerChain []HandleFunc, w http.ResponseWriter, r *http.Request) *Context {
	c := &Context{
		Request:        r,
		ResponseWriter: w,
		index:          -1,
		handlerChain:   handlerChain,
	}
	return c
}

func (t *Context) Next() {
	t.index += 1
	for t.index < len(t.handlerChain) {
		t.handlerChain[t.index](t)
		t.index += 1
	}
}

func (t *Context) Abort() {
	t.index = len(t.handlerChain) + 1
}

func (t *Context) WriteHeaderCode(code int) {
	t.StatusCode = code
	t.ResponseWriter.WriteHeader(code)
}

func (t *Context) Write(data []byte) (int, error) {
	return t.ResponseWriter.Write(data)
}

func (t *Context) WriteJsonWithCode(data interface{}, code int) (int, error) {
	t.ResponseWriter.Header().Set(HeaderContentType, "application/json;utf8")
	t.WriteHeaderCode(code)
	jsonData, _ := json.Marshal(data)
	return t.Write(jsonData)
}

func (t *Context) WriteJsonFailed(data interface{}) (int, error) {
	t.ResponseWriter.Header().Set(HeaderContentType, "application/json;utf8")
	t.WriteHeaderCode(http.StatusBadRequest)
	jsonData, _ := json.Marshal(data)
	return t.Write(jsonData)
}

func (t *Context) ERROR(data errors.Error) (int, error) {
	if t.IsProtoBuf() {
		return t.WriteProtobufFailed(dt.FromError(data))
	} else {
		return t.WriteJsonFailed(data)
	}
}

func (t *Context) SUCCESS(data interface{}) (int, error) {
	if t.IsProtoBuf() {
		resp, ok := data.(proto.Message)
		if !ok {
			panic("is not an instance of proto.Message")
		}
		return t.WriteProtobufSuccess(resp)
	} else {
		return t.WriteJsonSuccess(data)
	}
}

func (t *Context) DATA(data interface{}) (int, error) {
	if t.IsProtoBuf() {
		resp, ok := data.(proto.Message)
		if !ok {
			panic("is not an instance of proto.Message")
		}
		return t.WriteProtobufWithCode(resp, http.StatusOK)
	} else {
		return t.WriteJsonWithCode(data, http.StatusOK)
	}
}

func (t *Context) WriteJsonSuccess(data interface{}) (int, error) {
	t.ResponseWriter.Header().Set(HeaderContentType, "application/json;utf8")
	t.WriteHeaderCode(http.StatusOK)
	jsonData, _ := json.Marshal(JsonResponseWrapper{
		Code: 0,
		Data: data,
		Msg:  "",
	})
	return t.Write(jsonData)
}

func (t *Context) WriteProtobufWithCode(data proto.Message, code int) (int, error) {
	t.ResponseWriter.Header().Set(HeaderContentType, "application/protobuf")
	t.StatusCode = code
	t.WriteHeaderCode(code)
	buf, err := proto.Marshal(data)
	if err != nil {
		return 0, err
	}
	return t.Write(buf)
}

func (t *Context) WriteProtobufFailed(data proto.Message) (int, error) {
	return t.WriteProtobufWithCode(data, http.StatusBadRequest)
}

func (t *Context) WriteProtobufSuccess(data proto.Message) (int, error) {
	return t.WriteProtobufWithCode(data, http.StatusOK)
}

func (t *Context) GetHeader(k string) string {
	value := t.Request.Header.Get(k)
	if value == "" {
		value = t.Request.Header.Get(strings.ToLower(k))
	}
	return value
}

func (t *Context) GetContentType() string {
	return t.GetHeader("Content-Type")
}

func (t *Context) IsProtoBuf() bool {
	ct := t.GetContentType()
	ct = strings.ToLower(ct)
	if strings.Contains(ct, "application") && strings.Contains(ct, "protobuf") {
		return true
	}
	return false
}

func (t *Context) IsJson() bool {
	ct := t.GetContentType()
	ct = strings.ToLower(ct)
	if strings.Contains(ct, "application") && strings.Contains(ct, "json") {
		return true
	}
	return false
}

func (t *Context) GetCopyOfBodyBytes() ([]byte, error) {
	data, err := t.GetRequestBodyBytes()
	if err != nil {
		return nil, err
	}
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret, nil
}

func (t *Context) GetRequestBodyBytes() ([]byte, error) {
	if t.BodyContent != nil {
		return t.BodyContent, nil
	}
	data, err := ioutil.ReadAll(t.Request.Body)
	defer func() {
		_ = t.Request.Body.Close()
	}()
	if err != nil {
		return nil, err
	}
	t.BodyContent = data
	return t.BodyContent, nil
}

func (t *Context) IsPost() bool {
	return t.Request.Method == http.MethodPost
}

func (t *Context) GetToken() string {
	return t.GetHeader(yc.HeaderContextToken)
}

func (t *Context) GetSign() string {
	return t.GetHeader(yc.HeaderContextSign)
}

func (t *Context) GetSignMethod() yc.SignMethod {
	v := t.GetHeader(yc.HeaderContextSignMethod)
	v = strings.ToLower(v)
	if v == string(yc.SignMethodMd5) {
		return yc.SignMethodMd5
	} else if v == string(yc.SignMethodSha1) {
		return yc.SignMethodSha1
	} else if v == string(yc.SignMethodSha256) {
		return yc.SignMethodSha256
	} else if v == string(yc.SignMethodSha512) {
		return yc.SignMethodSha512
	}
	return yc.SignMethodUnknown
}

func (t *Context) GetAccessKey() string {
	return t.GetHeader(yc.HeaderContextAccessKey)
}

func (t *Context) GetTimestamp() int64 {
	v := t.GetHeader(yc.HeaderContextTimestamp)
	if v == "" {
		return 0
	}
	timestamp, _ := strconv.ParseInt(v, 10, 64)
	return timestamp
}

func (t *Context) SetUserId(userId *dt.ID) {
	t.RpcContext = metadata.NewIncomingContext(t.RpcContext, metadata.New(map[string]string{
		yc.MdContextUserId: fmt.Sprintf("%d", userId.Uint64),
	}))
}
