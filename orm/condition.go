package orm

import (
	"reflect"
	"strings"
)

type Condition interface {
	Clone() Condition
	Build() (statement string, values []interface{}, err error)
}

type expressionCondition struct {
	field    FieldName
	operator string
	value    interface{}
}

func (t expressionCondition) Clone() Condition {
	return expressionCondition{
		field:    t.field,
		operator: t.operator,
		value:    t.value,
	}
}

func NewExpressionCondition(field FieldName, operator string, value interface{}) Condition {
	return expressionCondition{
		field:    field,
		operator: operator,
		value:    value,
	}
}

func (t expressionCondition) Build() (statement string, values []interface{}, err error) {
	statement = t.field.Wrap().String() + wrap(t.operator, " ")
	value := reflect.ValueOf(t.value)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Slice {
		statement += "(" + strings.Trim(strings.Repeat("?, ", value.Len()), ", ") + ")"
		for i := 0; i < value.Len(); i++ {
			values = append(values, value.Index(i).Interface())
		}
	} else {
		statement += " ?"
		values = append(values, value.Interface())
	}
	return
}

type nullCondition struct {
	operator string
	field    FieldName
}

func (t nullCondition) Build() (statement string, values []interface{}, err error) {
	return t.field.Wrap().String() + " " + t.operator, nil, nil
}

func (t nullCondition) Clone() Condition {
	return nullCondition{
		operator: t.operator,
		field:    t.field,
	}
}

func NewNotNullCondition(field FieldName) Condition {
	return nullCondition{
		operator: "IS NOT NULL",
		field:    field,
	}
}

func NewNullCondition(field FieldName) Condition {
	return nullCondition{
		operator: "IS NULL",
		field:    field,
	}
}

type nestedCondition struct {
	whereWrapper WhereWrapper
}

func (t nestedCondition) Build() (statement string, values []interface{}, err error) {
	defer func() {
		statement = "(" + statement + ")"
	}()
	return t.whereWrapper.Build()
}

func NewNestedCondition(w WhereWrapper) Condition {
	return nestedCondition{
		whereWrapper: w,
	}
}

func (t nestedCondition) Clone() Condition {
	return nestedCondition{
		whereWrapper: t.whereWrapper.Clone(),
	}
}
