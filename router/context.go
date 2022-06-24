package router

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
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
	for t.index < len(t.handlerChain) {
		t.index += 1
		t.handlerChain[t.index](t)
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
	return t.WriteJsonWithCode(data, http.StatusBadRequest)
}

func (t *Context) WriteJsonSuccess(data interface{}) (int, error) {
	return t.WriteJsonWithCode(data, http.StatusOK)
}

func (t *Context) WriteProtobufWithCode(data proto.Message, code int) (int, error) {
	t.ResponseWriter.Header().Set(HeaderContentType, "application/protobuf")
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
