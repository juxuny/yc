package errors

import (
	"encoding/json"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"reflect"
	"strconv"
)

type Error struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (t Error) WithField(fieldName string, value interface{}) Error {
	ret := Error{
		Code: t.Code,
		Msg:  t.Msg,
		Data: map[string]interface{}{},
	}
	if t.Data != nil {
		for k, v := range t.Data {
			ret.Data[k] = v
		}
	}
	ret.Data[fieldName] = value
	return ret
}

func (t Error) Wrap(err error) Error {
	if err != nil {
		return t.WithField("prev", err.Error())
	}
	return t
}

func (t Error) SetMsg(msg string) Error {
	return Error{
		Code: t.Code,
		Msg:  msg,
	}
}

func (t Error) Err() error {
	if t.Code == 0 {
		return nil
	}
	return t
}

func (t Error) Error() string {
	jsonData, _ := json.Marshal(t)
	return string(jsonData)
}

func New(code int, msg string) error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func FromError(err error) (ret Error, ok bool) {
	if err == nil {
		return ret, true
	}
	s, ok := status.FromError(err)
	var jsonData string
	if ok {
		jsonData = s.Message()
	} else {
		jsonData = err.Error()
	}
	if err := json.Unmarshal([]byte(jsonData), &ret); err != nil {
		return ret, false
	}
	return ret, true
}

func InitErrorStruct(in interface{}) error {
	codeSet := make(map[int64]string)
	vv := reflect.ValueOf(in)
	if vv.Kind() != reflect.Ptr {
		return errors.Errorf("input must be a point of struct")
	}
	elem := vv.Elem()
	tt := elem.Type()
	for i := 0; i < tt.NumField(); i++ {
		if tt.Field(i).Type.String() != "errors.Error" {
			continue
		}
		var code int64
		var msg string
		var err error
		tag := tt.Field(i).Tag
		if v, ok := tag.Lookup("code"); !ok {
			return errors.Errorf("not found code in tag")
		} else {
			code, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
		}
		if v, ok := tag.Lookup("msg"); !ok {
			return errors.Errorf("not foun msg in tag")
		} else {
			msg = v
		}
		elem.Field(i).Set(reflect.ValueOf(Error{
			Code: int(code),
			Msg:  msg,
		}))
		if field, ok := codeSet[code]; ok {
			return errors.Errorf("fields '%s' and '%s' have duplicated code: %v", field, tt.Field(i).Name, code)
		} else {
			codeSet[code] = tt.Field(i).Name
		}
	}
	return nil
}
