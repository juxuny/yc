package template

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

func toUpperFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'a' && ret[0] <= 'z' {
		ret[0] = ret[0] - ('a' - 'A')
	}
	return string(ret)
}

func toLowerFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'A' && ret[0] <= 'Z' {
		ret[0] = ret[0] + ('a' - 'A')
	}
	return string(ret)
}

func trimPointer(s string) string {
	return strings.Trim(s, "*")
}

func raw(s interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("%v", s))
}

func castNumber(v interface{}) string {
	in := reflect.ValueOf(v)
	if !in.IsValid() {
		return ""
	}
	if in.Kind() == reflect.Ptr {
		if in.IsZero() {
			return ""
		}
		in = in.Elem()
	}
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", in.Convert(reflect.TypeOf(float64(0))).Interface())
	default:
	}
	return fmt.Sprintf("%v", in.Interface())
}
