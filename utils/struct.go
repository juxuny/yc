package utils

import (
	"github.com/juxuny/yc/errors"
	"reflect"
)

type structHelper struct{}

var StructHelper = &structHelper{}

func (structHelper) Patch(input interface{}, update map[string]interface{}) error {
	value := reflect.ValueOf(input)
	for k, v := range update {
		field := value.Elem().FieldByName(k)
		if !field.IsValid() {
			return errors.SystemError.ReflectNoFieldError.WithField("field", k)
		}
		field.Set(reflect.ValueOf(v))
	}
	return nil
}
