package orm

import (
	"database/sql"
	"fmt"
	"github.com/juxuny/yc/dt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type dataTypeConverter struct{}

var defaultConverter = &dataTypeConverter{}

func (t *dataTypeConverter) Convert(in reflect.Value, dstType reflect.Type) reflect.Value {
	if in.Kind() == reflect.Ptr {
		in = in.Elem()
	}
	var ret reflect.Value
	//fmt.Println("in: ", in.Type().String(), " to:", dstType.String(), " kind:", in.Kind())
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = t.convertInt(in, dstType)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret = t.convertUint(in, dstType)
	case reflect.Slice:
		ret = t.convertSlice(in, dstType)
	case reflect.Struct:
		ret = t.convertStruct(in, dstType)
	default:
		ret = reflect.Zero(dstType)
	}
	return ret
}

func (t *dataTypeConverter) convertStruct(in reflect.Value, dstType reflect.Type) reflect.Value {
	var ret reflect.Value
	if dstType.Kind() == reflect.Ptr {
		ret = reflect.New(dstType.Elem())
	} else {
		ret = reflect.New(dstType)
	}
	holder := ret.Elem()
	inTypeName := in.Type().String()
	holderTypeName := holder.Type().String()
	if inTypeName == "sql.NullInt64" {
		value := in.Interface().(sql.NullInt64)
		switch holder.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			holder.SetInt(value.Int64)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			holder.SetUint(uint64(value.Int64))
		case reflect.Float64, reflect.Float32:
			holder.SetFloat(float64(value.Int64))
		case reflect.String:
			holder.SetString(fmt.Sprintf("%v", value.Int64))
		case reflect.Bool:
			holder.SetBool(value.Int64 > 0)
		default:
			if holderTypeName == "sql.NullInt64" {
				holder.Set(in)
			} else if holderTypeName == "dt.NullInt64" {
				holder.Set(reflect.ValueOf(dt.NullInt64{
					Valid: value.Valid,
					Int64: value.Int64,
				}))
			} else if strings.Contains(holderTypeName, "dt.ID") {
				holder.Set(reflect.ValueOf(dt.ID{
					Valid:  value.Valid,
					Uint64: uint64(value.Int64),
				}))
			} else if strings.Contains(holderTypeName, "int") {
				if strings.Contains(holderTypeName, "*") {
					if !value.Valid {
						holder.Set(reflect.Zero(holder.Type()))
					} else {
						if !strings.Contains(holderTypeName, "uint") {
							v := reflect.New(holder.Type().Elem())
							v.Elem().SetInt(value.Int64)
							holder.Set(v)
						} else {
							v := reflect.New(holder.Type())
							v.Elem().SetUint(uint64(value.Int64))
							holder.Set(v.Elem())
						}
					}
				} else {
					if !strings.Contains(holderTypeName, "uint") {
						holder.SetInt(value.Int64)
					} else {
						holder.SetUint(uint64(value.Int64))
					}
				}
			} else {
				panic("unknown dest type: " + holderTypeName + " in type:" + inTypeName)
			}
		}
	} else if inTypeName == "sql.NullTime" {
		nullTime := in.Interface().(sql.NullTime)
		if holderTypeName == "sql.NullTime" {
			holder.Set(in)
		} else if holderTypeName == "*time.Time" {
			var v *time.Time
			if nullTime.Valid {
				holder.Set(reflect.ValueOf(&nullTime.Time))
			} else {
				holder.Set(reflect.ValueOf(v))
			}
		} else if holderTypeName == "time.Time" {
			holder.Set(reflect.ValueOf(nullTime.Time))
		} else if holderTypeName == "string" {
			holder.Set(reflect.ValueOf(nullTime.Time.String()))
		} else if holderTypeName == "dt.NullString" {
			holder.Set(reflect.ValueOf(dt.NullString{
				Valid:   nullTime.Valid,
				String_: nullTime.Time.String(),
			}))
		} else {
			panic("unknown dest type: " + holderTypeName)
		}
	} else if inTypeName == "sql.NullFloat64" {
		nullFloat64 := in.Interface().(sql.NullFloat64)
		if strings.Contains(holderTypeName, "float") {
			if strings.Contains(holderTypeName, "*") {
				if !nullFloat64.Valid {
					var v *float64
					holder.Set(reflect.ValueOf(v))
				} else {
					holder.Set(reflect.ValueOf(&nullFloat64.Float64))
				}
			} else {
				holder.SetFloat(nullFloat64.Float64)
			}
		} else if holderTypeName == "dt.NullFloat64" {
			holder.Set(reflect.ValueOf(dt.NullFloat64{
				Valid:   nullFloat64.Valid,
				Float64: nullFloat64.Float64,
			}))
		} else if holderTypeName == "dt.NullFloat32" {
			holder.Set(reflect.ValueOf(dt.NullFloat32{
				Valid:   nullFloat64.Valid,
				Float32: float32(nullFloat64.Float64),
			}))
		} else if holderTypeName == "sql.NullFloat64" {
			holder.Set(in)
		} else {
			panic("unknown struct: " + inTypeName)
		}
	} else if inTypeName == "sql.NullBool" {
		nullBool := in.Interface().(sql.NullBool)
		if holderTypeName == "sql.NullBool" {
			holder.Set(in)
		} else if holderTypeName == "dt.NullBool" {
			holder.Set(reflect.ValueOf(dt.NullBool{
				Valid: nullBool.Valid,
				Bool:  nullBool.Bool,
			}))
		} else if holderTypeName == "bool" {
			holder.SetBool(nullBool.Bool)
		} else if holderTypeName == "*bool" {
			holder.Set(reflect.ValueOf(&nullBool.Bool))
		} else if holderTypeName == "int" || holderTypeName == "int8" || holderTypeName == "int16" || holderTypeName == "int32" || holderTypeName == "int64" || holderTypeName == "uint" || holderTypeName == "uint8" || holderTypeName == "uint16" || holderTypeName == "uint32" || holderTypeName == "uint64" {
			var v int
			if nullBool.Bool {
				v = 1
			}
			holder.Set(reflect.ValueOf(v).Convert(holder.Type()))
		} else {
			panic("unknown dest type:" + holderTypeName)
		}
	} else {
		panic("unknown struct: " + inTypeName)
	}
	if dstType.Kind() == reflect.Ptr {
		return ret
	} else {
		return ret.Elem()
	}
}

func (t *dataTypeConverter) convertUint(in reflect.Value, dstType reflect.Type) reflect.Value {
	var ret reflect.Value
	if dstType.Kind() == reflect.Ptr {
		ret = reflect.New(dstType.Elem())
	} else {
		ret = reflect.New(dstType)
	}
	holder := ret.Elem()
	switch holder.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		holder.Set(in.Convert(dstType))
	case reflect.Struct:
		if holder.Type().String() == "dt.ID" {
			holder.Set(reflect.ValueOf(dt.NewID(in.Uint())))
		} else {
			panic(fmt.Errorf("unknown dstType: %v", dstType.String()))
		}
	}
	if dstType.Kind() == reflect.Ptr {
		return ret
	} else {
		return ret.Elem()
	}
}

func (t *dataTypeConverter) convertInt(in reflect.Value, dstType reflect.Type) reflect.Value {
	var ret reflect.Value
	if dstType.Kind() == reflect.Ptr {
		ret = reflect.New(dstType.Elem())
	} else {
		ret = reflect.New(dstType)
	}
	holder := ret.Elem()
	switch holder.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		holder.Set(in.Convert(dstType))
	case reflect.Struct:
		if holder.Type().String() == "dt.ID" {
			holder.Set(reflect.ValueOf(dt.NewID(uint64(in.Int()))))
		} else {
			panic("unknown dest type: " + dstType.String())
		}
	case reflect.Bool:
		holder.SetBool(in.Int() > 0)
	default:
		panic("unknown dest type: " + dstType.String())
	}
	if dstType.Kind() == reflect.Ptr {
		return ret
	} else {
		return ret.Elem()
	}
}

func (t *dataTypeConverter) convertSlice(in reflect.Value, dstType reflect.Type) reflect.Value {
	var ret reflect.Value
	if dstType.Kind() == reflect.Ptr {
		ret = reflect.New(dstType.Elem())
	} else {
		ret = reflect.New(dstType)
	}
	holder := ret.Elem()
	inTypeName := in.Type().String()
	holderTypeName := holder.Type().String()
	if inTypeName == "sql.RawBytes" {
		data := in.Interface().(sql.RawBytes)
		if holderTypeName == "sql.NullString" {
			holder.Set(reflect.ValueOf(sql.NullString{
				Valid:  data != nil,
				String: string(data),
			}))
		} else if holderTypeName == "string" {
			holder.SetString(string(data))
		} else if holderTypeName == "dt.NullString" {
			holder.Set(reflect.ValueOf(dt.NullString{
				Valid:   data != nil,
				String_: string(data),
			}))
		} else if holderTypeName == "sql.NullTime" {
			holder.Set(reflect.Zero(dstType))
		} else if holderTypeName == "*time.Time" {
			holder.Set(reflect.Zero(dstType))
		} else if strings.Contains(holderTypeName, "float") {
			v, _ := strconv.ParseFloat(string(data), 64)
			if strings.Contains(holderTypeName, "*") {
				holder.Set(reflect.ValueOf(&v))
			} else {
				holder.SetFloat(v)
			}
		} else {
			panic("unknown dest type: " + holderTypeName)
		}
	} else {
		panic("unknown in type: " + inTypeName)
	}
	if dstType.Kind() == reflect.Ptr {
		return ret
	} else {
		return ret.Elem()
	}
}

func (t *dataTypeConverter) convertToWhereHolderInterface(in reflect.Value) interface{} {
	switch in.Kind() {
	case reflect.Ptr:
		if in.IsNil() {
			return nil
		}
		inType := in.Type().String()
		if inType == "*dt.ID" {
			data := in.Interface().(*dt.ID)
			if data.Valid {
				return data.Uint64
			} else {
				return nil
			}
		} else if inType == "*dt.NullString" {
			data := in.Interface().(*dt.NullString)
			if data.Valid {
				return data.String_
			} else {
				return nil
			}
		} else if inType == "*dt.NullBool" {
			data := in.Interface().(*dt.NullBool)
			if data.Valid {
				return data.Bool
			} else {
				return nil
			}
		} else if inType == "*dt.NullInt64" {
			data := in.Interface().(*dt.NullInt64)
			if data.Valid {
				return data.Int64
			} else {
				return nil
			}
		} else if inType == "*dt.NullInt32" {
			data := in.Interface().(*dt.NullInt32)
			if data.Valid {
				return data.Int32
			} else {
				return nil
			}
		} else if inType == "*dt.NullFloat64" {
			data := in.Interface().(*dt.NullFloat64)
			if data.Valid {
				return data.Float64
			} else {
				return nil
			}
		} else if inType == "*dt.NullFloat32" {
			data := in.Interface().(*dt.NullFloat32)
			if data.Valid {
				return data.Float32
			} else {
				return nil
			}
		} else {
			return t.convertToWhereHolderInterface(in.Elem())
		}
	default:
		return in.Interface()
	}
}
