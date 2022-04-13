package validator

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"reflect"
	"strconv"
)

type lengthMaxValidator struct{}

func (t *lengthMaxValidator) Run(v interface{}, refValueString string) (bool, error) {
	vv := reflect.ValueOf(v)
	tt := vv.Type()
	lengthOfRefValue, err := strconv.ParseInt(refValueString, 10, 64)
	if err != nil {
		return false, errors.SystemError.InvalidRefValueDefinition.Wrap(err)
	}
	if tt.Kind() == reflect.Slice {
		return vv.Len() <= int(lengthOfRefValue), nil
	} else if tt.Kind() == reflect.String {
		return len(vv.String()) <= int(lengthOfRefValue), nil
	} else {
		return false, errors.SystemError.InvalidDataType.Wrap(fmt.Errorf("kind: %v", tt.Kind()))
	}
}

type lengthMinValidator struct{}

func (t *lengthMinValidator) Run(v interface{}, refValueString string) (bool, error) {
	vv := reflect.ValueOf(v)
	tt := vv.Type()
	lengthOfRefValue, err := strconv.ParseInt(refValueString, 10, 64)
	if err != nil {
		return false, errors.SystemError.InvalidRefValueDefinition.Wrap(err)
	}
	if tt.Kind() == reflect.Slice {
		return vv.Len() >= int(lengthOfRefValue), nil
	} else if tt.Kind() == reflect.String {
		return len(vv.String()) >= int(lengthOfRefValue), nil
	} else {
		return false, errors.SystemError.InvalidDataType.Wrap(fmt.Errorf("kind: %v", tt.Kind()))
	}
}
