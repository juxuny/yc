package orm

import "reflect"

type UpdateExpression interface {
	Build() (statement string, values []interface{}, err error)
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

type setValueExpression struct {
	field FieldName
	value interface{}
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
