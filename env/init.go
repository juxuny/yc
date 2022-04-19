package env

import (
	"fmt"
	"github.com/fatih/camelcase"
	"reflect"
	"strconv"
	"strings"
)

func Init(ks interface{}, upper bool, prefix ...string) {
	v := reflect.ValueOf(ks)
	tt := reflect.TypeOf(ks)
	tt = tt.Elem()
	p := ""
	if len(prefix) > 0 {
		p = prefix[0]
	}
	if p != "" && !strings.Contains(p, "_") {
		p += "_"
	}
	for i := 0; i < tt.NumField(); i++ {
		value := strings.ToLower(strings.Join(camelcase.Split(tt.Field(i).Name), "_"))
		if upper {
			value = strings.ToUpper(value)
		}
		value = strings.Trim(value, "_")
		if p != "" {
			value = p + value
		}
		defaultValue, ok := tt.Field(i).Tag.Lookup("env")
		var finalValue string
		if ok {
			finalValue = GetString(value, defaultValue)
		} else {
			finalValue = GetString(value)
		}
		switch tt.Field(i).Type.Kind() {
		case reflect.String:
			v.Elem().Field(i).SetString(finalValue)
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			intValue, err := strconv.ParseInt(finalValue, 10, 64)
			if err != nil {
				panic(fmt.Errorf("%s is not a int: \"%v\"", value, finalValue))
			}
			v.Elem().Field(i).SetInt(intValue)
		case reflect.Uint64, reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			intValue, err := strconv.ParseUint(finalValue, 10, 64)
			if err != nil {
				panic(fmt.Errorf("%s is not a uint: \"%v\"", value, finalValue))
			}
			v.Elem().Field(i).SetUint(intValue)
		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(finalValue, 64)
			if err != nil {
				panic(fmt.Errorf("%s is not a float64: \"%v\"", value, finalValue))
			}
			v.Elem().Field(i).SetFloat(floatValue)
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(finalValue)
			if err != nil {
				panic(fmt.Errorf("%s is not a bool: \"%v\"", value, finalValue))
			}
			v.Elem().Field(i).SetBool(boolValue)
		default:
			panic("unknown type: " + tt.Field(i).Type.Name())
		}
	}
}
