package cos

import (
	"reflect"
	"strconv"
)
import "github.com/pkg/errors"

func Parse(data map[string]string, out interface{}) error {
	vv := reflect.ValueOf(out)
	vv = vv.Elem()
	tt := vv.Type()
	for configKey, value := range data {
		for i := 0; i < tt.NumField(); i++ {
			f := tt.Field(i)
			if tagValue, ok := f.Tag.Lookup("cos"); ok {
				if tagValue != configKey {
					continue
				}
				kind := vv.Field(i).Kind()
				if kind == reflect.String {
					vv.Field(i).SetString(value)
				} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
					v, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					vv.Field(i).SetInt(v)
				} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
					v, err := strconv.ParseUint(value, 10, 64)
					if err != nil {
						return err
					}
					vv.Field(i).SetUint(v)
				} else if kind == reflect.Float64 || kind == reflect.Float32 {
					v, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
					vv.Field(i).SetFloat(v)
				} else if kind == reflect.Bool {
					v, err := strconv.ParseBool(value)
					if err != nil {
						return err
					}
					vv.Field(i).SetBool(v)
				} else {
					return errors.New("unknown kind of field: " + f.Name)
				}
			}
		}
	}
	return nil
}
