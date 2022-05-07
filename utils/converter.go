package utils

import (
	"database/sql"
	"fmt"
	"github.com/juxuny/yc/dt"
	"reflect"
	"strconv"
)

type converter struct{}

var Converter = converter{}

func (converter) ToFloat64(v interface{}) (float64, error) {
	in := reflect.ValueOf(v)
	if !in.IsValid() {
		return 0, fmt.Errorf("invalid value: %v", v)
	}
	if in.Kind() == reflect.Ptr {
		in = in.Elem()
	}
	inType := in.Type().String()
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return in.Convert(reflect.TypeOf(float64(0))).Float(), nil
	case reflect.String:
		return strconv.ParseFloat(in.String(), 10)
	case reflect.Struct:
		if inType == "dt.ID" {
			v := in.Interface().(dt.ID)
			return float64(v.Uint64), nil
		} else if inType == "dt.NullString" {
			return strconv.ParseFloat(in.Interface().(dt.NullString).String_, 10)
		} else if inType == "dt.NullInt64" {
			return float64(in.Interface().(dt.NullInt64).Int64), nil
		} else if inType == "dt.NullInt32" {
			return float64(in.Interface().(dt.NullInt32).Int32), nil
		} else if inType == "dt.NullBool" {
			v := in.Interface().(dt.NullBool)
			if v.Bool {
				return 1, nil
			} else {
				return 0, nil
			}
		} else if inType == "dt.NullFloat64" {
			return in.Interface().(dt.NullFloat64).Float64, nil
		} else if inType == "dt.NullFloat32" {
			return float64(in.Interface().(dt.NullFloat32).Float32), nil
		} else if inType == "sql.NullString" {
			return strconv.ParseFloat(in.Interface().(sql.NullString).String, 10)
		} else if inType == "sql.NullInt64" {
			return float64(in.Interface().(sql.NullInt64).Int64), nil
		} else if inType == "sql.NullFloat64" {
			return in.Interface().(sql.NullFloat64).Float64, nil
		} else if inType == "sql.NullBool" {
			v := in.Interface().(sql.NullBool)
			if v.Bool {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	default:
	}
	return 0, fmt.Errorf("unknown data type: %v", in.Type().String())
}
