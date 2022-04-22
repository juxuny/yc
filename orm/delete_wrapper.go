package orm

import "fmt"

type DeleteWrapper interface {
	Model(model interface{}) DeleteWrapper
	WhereWrapper
	TableName(tableName TableName) DeleteWrapper
	Build() (statement string, values []interface{}, err error)
}

type deleteWrapper struct {
	model     Model
	tableName TableName
	WhereWrapper
}

func (t *deleteWrapper) Model(model interface{}) DeleteWrapper {
	t.model = CreateModel(model)
	return t
}

func (t *deleteWrapper) TableName(tableName TableName) DeleteWrapper {
	t.tableName = tableName
	return t
}

func (t *deleteWrapper) Build() (statement string, values []interface{}, err error) {
	if t.tableName == "" {
		t.tableName = t.model.TableName
	}
	if t.tableName == "" {
		return statement, values, fmt.Errorf("missing table_name")
	}
	statement = "DELETE FROM " + t.tableName.Wrap().String()
	if t.WhereWrapper == nil {
		return statement, values, fmt.Errorf("missing where statement")
	}
	s, vs, err := t.WhereWrapper.Build()
	if err != nil {
		return statement, values, err
	}
	if s == "" {
		return statement, values, fmt.Errorf("missing where statement")
	}
	statement += " WHERE " + s
	values = append(values, placementValueFilter(vs)...)
	return
}

func NewDeleteWrapper(model ...interface{}) DeleteWrapper {
	w := &deleteWrapper{
		WhereWrapper: NewAndWhereWrapper(),
	}
	if len(model) > 0 {
		return w.Model(model[0])
	}
	return w
}
