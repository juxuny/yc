package validator

import (
	"reflect"
)

type requiredValidator struct{}

func (t *requiredValidator) Run(v interface{}, refValueString string) (bool, error) {
	vv := reflect.ValueOf(v)
	if vv.IsZero() && (refValueString == "true" || refValueString == "") {
		return false, nil
	}
	return true, nil
}
