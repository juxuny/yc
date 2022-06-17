package orm

import (
	"database/sql"
	"fmt"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"reflect"
	"strings"
	"time"
)

func generateSlotFromColumnTypes(columnTypes []*sql.ColumnType) (values []reflect.Value, holder []interface{}) {
	for _, col := range columnTypes {
		v := reflect.New(col.ScanType())
		values = append(values, v)
		holder = append(holder, v.Interface())
	}
	return
}

func getOrmColumnNameMap(model reflect.Type) map[string]string {
	ret := make(map[string]string)
	for i := 0; i < model.NumField(); i++ {
		ft := model.Field(i)
		tag := ft.Tag
		if v, b := tag.Lookup("orm"); b {
			ret[strings.TrimSpace(v)] = ft.Name
		}
	}
	return ret
}

func wrap(s string, c string) string {
	return c + s + c
}

func JoinFieldNames(fields []FieldName, sep string) string {
	list := make([]string, len(fields))
	for i, item := range fields {
		list[i] = item.String()
	}
	return strings.Join(list, sep)
}

func placementValueFilter(values []interface{}) []interface{} {
	ret := make([]interface{}, len(values))
	for i, item := range values {
		ret[i] = toPlacementValue(reflect.ValueOf(item))
	}
	return ret
}

func toPlacementValue(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool, reflect.Float32, reflect.Float64:
		return v.Interface()
	case reflect.Struct:
		vn := v.Type().String()
		if vn == "sql.NullInt64" {
			data := v.Interface().(sql.NullInt64)
			if data.Valid {
				return data.Int64
			} else {
				return nil
			}
		} else if vn == "sql.NullInt32" {
			data := v.Interface().(sql.NullInt32)
			if data.Valid {
				return data.Int32
			} else {
				return nil
			}
		} else if vn == "sql.NullBool" {
			data := v.Interface().(sql.NullBool)
			if data.Valid {
				return data.Bool
			} else {
				return nil
			}
		} else if vn == "sql.NullString" {
			data := v.Interface().(sql.NullString)
			if data.Valid {
				return data.String
			} else {
				return nil
			}
		} else if vn == "sql.NullFloat64" {
			data := v.Interface().(sql.NullFloat64)
			if data.Valid {
				return data.Float64
			} else {
				return nil
			}
		} else if vn == "sql.NullTime" {
			data := v.Interface().(sql.NullTime)
			if data.Valid {
				return data.Time
			} else {
				return nil
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		vn := v.Type().String()
		if vn == "*dt.ID" {
			data := v.Interface().(*dt.ID)
			if data.Valid {
				return data.Uint64
			} else {
				return nil
			}
		} else if vn == "*dt.NullInt64" {
			data := v.Interface().(*dt.NullInt64)
			if data.Valid {
				return data.Int64
			} else {
				return nil
			}
		} else if vn == "*dt.NullInt32" {
			data := v.Interface().(*dt.NullInt32)
			if data.Valid {
				return data.Int32
			} else {
				return nil
			}
		} else if vn == "*dt.NullString" {
			data := v.Interface().(*dt.NullString)
			if data.Valid {
				return data.String_
			} else {
				return nil
			}
		} else if vn == "*dt.NullBool" {
			data := v.Interface().(*dt.NullBool)
			if data.Valid {
				return data.Bool
			} else {
				return nil
			}
		} else if vn == "*dt.NullFloat64" {
			data := v.Interface().(*dt.NullFloat64)
			if data.Valid {
				return data.Float64
			} else {
				return nil
			}
		} else if vn == "*dt.NullFloat32" {
			data := v.Interface().(*dt.NullFloat32)
			if data.Valid {
				return data.Float32
			} else {
				return nil
			}
		} else if vn == "*sql.NullInt64" {
			data := v.Interface().(*sql.NullInt64)
			if data.Valid {
				return data.Int64
			} else {
				return nil
			}
		} else if vn == "*sql.NullInt32" {
			data := v.Interface().(*sql.NullInt32)
			if data.Valid {
				return data.Int32
			} else {
				return nil
			}
		} else if vn == "*sql.NullBool" {
			data := v.Interface().(*sql.NullBool)
			if data.Valid {
				return data.Bool
			} else {
				return nil
			}
		} else if vn == "*sql.NullString" {
			data := v.Interface().(*sql.NullString)
			if data.Valid {
				return data.String
			} else {
				return nil
			}
		} else if vn == "*sql.NullFloat64" {
			data := v.Interface().(*sql.NullFloat64)
			if data.Valid {
				return data.Float64
			} else {
				return nil
			}
		} else if vn == "*sql.NullTime" {
			data := v.Interface().(*sql.NullTime)
			if data.Valid {
				return data.Time
			} else {
				return nil
			}
		}
		return toPlacementValue(v.Elem())
	}
	if v.IsValid() {
		return v.Interface()
	}
	return nil
}

func printSqlStatement(statement string, values ...interface{}) {
	s := strings.ReplaceAll(statement, "?", "%v")
	l := make([]interface{}, len(values))
	for i, item := range values {
		vv := reflect.ValueOf(item)
		if !vv.IsValid() {
			l[i] = "NULL"
			continue
		}
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			l[i] = fmt.Sprintf("%v", vv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			l[i] = fmt.Sprintf("%v", vv.Uint())
		case reflect.Float64, reflect.Float32:
			l[i] = fmt.Sprintf("%v", vv.Float())
		case reflect.String:
			l[i] = fmt.Sprintf("'%s'", vv.String())
		case reflect.Bool:
			l[i] = fmt.Sprintf("%v", vv.Bool())
		default:
			tn := vv.Type().String()
			if item == nil {
				l[i] = "NULL"
			} else if tn == "time.Time" {
				t := item.(time.Time)
				l[i] = wrap(t.Format("2006-01-02 15:04:05.000"), "'")
			} else if tn == "*time.Time" {
				t := item.(*time.Time)
				l[i] = wrap(t.Format("2006-01-02 15:04:05.000"), "'")
			} else {
				l[i] = fmt.Sprintf("'%v'", item)
			}
		}
	}
	s = fmt.Sprintf(s, l...)
	log.Info(s)
}

func cloneValue(in reflect.Value) reflect.Value {
	inputType := in.Type().String()
	if inputType == "*sql.RawBytes" {
		data, _ := in.Interface().(*sql.RawBytes)
		ret := make(sql.RawBytes, len(*data))
		copy(ret, *data)
		return reflect.ValueOf(&ret)
	} else if inputType == "sql.RawBytes" {
		data, _ := in.Interface().(sql.RawBytes)
		ret := make(sql.RawBytes, len(data))
		copy(ret, data)
		return reflect.ValueOf(ret)
	}
	return in
}
