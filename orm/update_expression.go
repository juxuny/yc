package orm

import (
	"fmt"
	"github.com/juxuny/yc/utils"
	"reflect"
)

type UpdateExpression interface {
	Build() (statement string, values []interface{}, err error)
	Clone() UpdateExpression
}

type duplicatedUpdateExpression struct {
	fields []FieldName
}

func (t duplicatedUpdateExpression) Build() (statement string, values []interface{}, err error) {
	for _, f := range t.fields {
		if statement != "" {
			statement += ", "
		}
		statement += f.Wrap().String() + " = " + "VALUES(" + f.Wrap().String() + ")"
	}
	return
}

func NewDuplicatedUpdateExpression(fields ...FieldName) UpdateExpression {
	u := duplicatedUpdateExpression{
		fields: fields,
	}
	return u
}

func (t duplicatedUpdateExpression) Clone() UpdateExpression {
	return duplicatedUpdateExpression{
		fields: append([]FieldName{}, t.fields...),
	}
}

type setValueExpression struct {
	field FieldName
	value interface{}
}

func (t setValueExpression) Clone() UpdateExpression {
	return setValueExpression{
		field: t.field,
		value: t.value,
	}
}

func (t setValueExpression) Build() (statement string, values []interface{}, err error) {
	statement = t.field.Wrap().String() + " = ?"
	values = []interface{}{toPlacementValue(reflect.ValueOf(t.value))}
	return
}

func NewSetValueUpdateExpression(field FieldName, value interface{}) UpdateExpression {
	expr := setValueExpression{
		field: field,
		value: value,
	}
	return expr
}

type increaseExpression struct {
	field FieldName
	value interface{}
}

func (t increaseExpression) Build() (statement string, values []interface{}, err error) {
	statement = t.field.Wrap().String() + " = " + t.field.Wrap().String() + " + ?"
	values = append(values, t.value)
	return
}

func (t increaseExpression) Clone() UpdateExpression {
	return increaseExpression{
		field: t.field,
		value: t.value,
	}
}

func NewIncreaseExpression(field FieldName, value interface{}) UpdateExpression {
	return &increaseExpression{
		field: field,
		value: value,
	}
}

type updateExpression struct {
	modelOrMap     interface{}
	ignoreFieldMap map[FieldName]struct{}
}

func (t updateExpression) Clone() UpdateExpression {
	ignoreFieldMap := make(map[FieldName]struct{})
	for k := range t.ignoreFieldMap {
		ignoreFieldMap[k] = struct{}{}
	}
	return updateExpression{
		modelOrMap:     t.modelOrMap,
		ignoreFieldMap: ignoreFieldMap,
	}
}

func (t updateExpression) Build() (statement string, values []interface{}, err error) {
	vv := reflect.ValueOf(t.modelOrMap)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	if vv.Kind() == reflect.Map {
		iterator := vv.MapRange()
		for iterator.Next() {
			if statement != "" {
				statement += ", "
			}
			k := iterator.Key()
			v := iterator.Value()
			if k.Kind() != reflect.String {
				return statement, values, fmt.Errorf("unknown key type: %v", k.Type().String())
			}
			statement += " " + FieldName(k.String()).Wrap().String() + " = ?"
			values = append(values, v.Interface())
		}
	} else if vv.Kind() == reflect.Struct {
		vt := vv.Type()
		for i := 0; i < vv.NumField(); i++ {
			ft := vt.Field(i)
			fv := vv.Field(i)
			fn := utils.ToUnderLine(ft.Name)
			if n, ok := ft.Tag.Lookup("orm"); ok {
				fn = n
			}
			if _, ok := ft.Tag.Lookup("ignore"); ok {
				continue
			}
			if statement != "" {
				statement += ", "
			}
			statement += FieldName(fn).Wrap().String() + " = ?"
			values = append(values, fv.Interface())
		}
	} else {
		return statement, values, fmt.Errorf("unknown data type: %v", vv.Type().String())
	}
	return
}

func NewUpdateExpression(modelOrMap interface{}, ignoreFields ...FieldName) UpdateExpression {
	ignoreFieldMap := make(map[FieldName]struct{})
	for _, item := range ignoreFields {
		ignoreFieldMap[item] = struct{}{}
	}
	return updateExpression{
		modelOrMap: modelOrMap,
	}
}
