package orm

import (
	"fmt"
	"github.com/juxuny/yc/log"
	"reflect"
)

type WhereWrapper interface {
	Eq(field FieldName, v interface{}) WhereWrapper
	Neq(field FieldName, v interface{}) WhereWrapper
	Le(field FieldName, v interface{}) WhereWrapper
	Lt(field FieldName, v interface{}) WhereWrapper
	Gt(field FieldName, v interface{}) WhereWrapper
	Ge(field FieldName, v interface{}) WhereWrapper
	In(field FieldName, v interface{}) WhereWrapper
	NotIn(field FieldName, v interface{}) WhereWrapper
	IsNotNull(field FieldName) WhereWrapper
	IsNull(field FieldName) WhereWrapper
	Nested(w WhereWrapper) WhereWrapper
	Like(field FieldName, v interface{}) WhereWrapper
	Clone() WhereWrapper
	Build() (string, []interface{}, error)
}

type ConditionLogic string

const (
	ConditionLogicAnd = ConditionLogic("AND")
	ConditionLogicOr  = ConditionLogic("OR")
)

type whereWrapper struct {
	logic      ConditionLogic
	conditions []Condition
}

func (t *whereWrapper) Clone() WhereWrapper {
	conditions := make([]Condition, 0)
	for _, item := range t.conditions {
		conditions = append(conditions, item.Clone())
	}
	return &whereWrapper{
		logic:      t.logic,
		conditions: append([]Condition{}, t.conditions...),
	}
}

func (t *whereWrapper) Like(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "LIKE", v))
	return t
}

func (t *whereWrapper) Nested(w WhereWrapper) WhereWrapper {
	t.conditions = append(t.conditions, NewNestedCondition(w))
	return t
}

func (t *whereWrapper) IsNotNull(field FieldName) WhereWrapper {
	t.conditions = append(t.conditions, NewNotNullCondition(field))
	return t
}

func (t *whereWrapper) IsNull(field FieldName) WhereWrapper {
	t.conditions = append(t.conditions, NewNullCondition(field))
	return t
}

func (t *whereWrapper) Eq(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "=", v))
	return t
}

func (t *whereWrapper) Neq(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "<>", v))
	return t
}

func (t *whereWrapper) Le(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "<=", v))
	return t
}

func (t *whereWrapper) Lt(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "<", v))
	return t
}

func (t *whereWrapper) Gt(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, ">", v))
	return t
}

func (t *whereWrapper) Ge(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, ">=", v))
	return t
}

func (t *whereWrapper) In(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "IN", v))
	return t
}

func (t *whereWrapper) NotIn(field FieldName, v interface{}) WhereWrapper {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr && vv.IsNil() {
		log.Debug("ignore nil: ", field)
		return t
	}
	t.conditions = append(t.conditions, NewExpressionCondition(field, "NOT IN", v))
	return t
}

func (t *whereWrapper) Build() (statement string, values []interface{}, err error) {
	for _, c := range t.conditions {
		if statement != "" {
			statement += " " + string(t.logic) + " "
		}
		if s, vs, err := c.Build(); err != nil {
			return "", nil, fmt.Errorf("syntax error: %v", err)
		} else {
			statement += s
			values = append(values, vs...)
		}
	}
	return
}

func NewAndWhereWrapper() WhereWrapper {
	return &whereWrapper{
		logic: ConditionLogicAnd,
	}
}

func NewOrWhereWrapper() WhereWrapper {
	return &whereWrapper{
		logic: ConditionLogicOr,
	}
}
