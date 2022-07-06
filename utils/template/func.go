package template

import (
	"fmt"
	"github.com/fatih/camelcase"
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

func lower(s string) string {
	return strings.ToLower(s)
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func splitVariable(v string) []string {
	if strings.Contains(v, "_") {
		return strings.Split(v, "_")
	} else if strings.Contains(v, "-") {
		return strings.Split(v, "-")
	} else {
		return camelcase.Split(v)
	}
}

func underlineString(v string) string {
	l := splitVariable(v)
	return strings.Join(l, "_")
}

func camelcaseString(v string) string {
	s := splitVariable(v)
	for i := range s {
		s[i] = strings.ToLower(s[i])
		s[i] = toUpperFirst(s[i])
	}
	return strings.Join(s, "")
}
