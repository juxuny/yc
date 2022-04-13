package validator

import (
	"fmt"
	"github.com/juxuny/yc/utils"
	"reflect"
	"strings"
)

type inValidator struct {
}

func (t *inValidator) Run(v interface{}, refValueString string) (bool, error) {
	vv := reflect.ValueOf(v)
	var inputString string
	switch vv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		inputString = fmt.Sprintf("%d", v)
	case reflect.Float32, reflect.Float64:
		inputString = fmt.Sprintf("%f", v)
	default:
		inputString = fmt.Sprintf("%v", v)
	}
	inputString = strings.TrimSpace(inputString)
	items := strings.Split(refValueString, ",")
	items = utils.StringHelper.Transform(items, strings.TrimSpace)
	refSet := utils.NewStringSet(items...)
	return refSet.Exists(inputString), nil
}
