package orm

import "fmt"

type UpdateWrapper interface {
	Model(v interface{}) UpdateWrapper
	WhereWrapper
	LeftJoin(tableName TableName, where WhereWrapper) UpdateWrapper
	RightJoin(tableName TableName, where WhereWrapper) UpdateWrapper
	SetValue(field FieldName, value interface{}) UpdateWrapper
	SetExpression(expressions ...UpdateExpression) UpdateWrapper
	Inc(field FieldName, value interface{}) UpdateWrapper
	Updates(object interface{}, ignoreFields ...FieldName) UpdateWrapper
	Build() (statement string, values []interface{}, err error)
	Table(tableName TableName) UpdateWrapper
	SetWhere(w WhereWrapper) UpdateWrapper
}

type updateWrapper struct {
	model     Model
	tableName TableName
	WhereWrapper
	joinList    []Join
	expressions []UpdateExpression
}

func (t *updateWrapper) SetWhere(w WhereWrapper) UpdateWrapper {
	t.WhereWrapper = w
	return t
}

func (t *updateWrapper) Inc(field FieldName, value interface{}) UpdateWrapper {
	t.expressions = append(t.expressions, NewIncreaseExpression(field, value))
	return t
}

func (t *updateWrapper) Updates(modelOrMap interface{}, ignoreFields ...FieldName) UpdateWrapper {
	t.expressions = append(t.expressions, NewUpdateExpression(modelOrMap, ignoreFields...))
	return t
}

func (t *updateWrapper) Table(tableName TableName) UpdateWrapper {
	t.tableName = tableName
	return t
}

func (t *updateWrapper) Model(v interface{}) UpdateWrapper {
	t.model = CreateModel(v)
	return t
}

func (t *updateWrapper) LeftJoin(tableName TableName, where WhereWrapper) UpdateWrapper {
	t.joinList = append(t.joinList, Join{
		Statement:    JoinStatementLeftJoin,
		TableName:    tableName,
		WhereWrapper: where,
	})
	return t
}

func (t *updateWrapper) RightJoin(tableName TableName, where WhereWrapper) UpdateWrapper {
	t.joinList = append(t.joinList, Join{
		Statement:    JoinStatementRightJoin,
		TableName:    tableName,
		WhereWrapper: where,
	})
	return t
}

func (t *updateWrapper) SetValue(field FieldName, value interface{}) UpdateWrapper {
	t.expressions = append(t.expressions, NewSetValueUpdateExpression(field, value))
	return t
}

func (t *updateWrapper) SetExpression(expressions ...UpdateExpression) UpdateWrapper {
	t.expressions = append(t.expressions, expressions...)
	return t
}

func (t *updateWrapper) Build() (statement string, values []interface{}, err error) {
	statement = "UPDATE"
	if t.tableName == "" {
		t.tableName = t.model.TableName.Wrap()
	}
	if t.tableName == "" {
		return "", nil, fmt.Errorf("missing table name")
	}
	statement += " " + t.tableName.String()
	joinStatement := ""
	if len(t.joinList) > 0 {
		for _, j := range t.joinList {
			if joinStatement != "" {
				joinStatement += " "
			}
			joinStatement += " " + string(j.Statement) + " " + j.TableName.Wrap().String()
			if j.WhereWrapper != nil {
				where, whereValueHolder, err := j.WhereWrapper.Build()
				if err != nil {
					return "", nil, fmt.Errorf("syntax error: %v", err)
				}
				joinStatement += " ON " + where
				values = append(values, whereValueHolder...)
			}
		}
	}
	if len(t.expressions) == 0 {
		return statement, values, fmt.Errorf("missing set statement")
	}
	statement += " SET "
	expressionStatement := ""
	expressionValues := make([]interface{}, 0)
	for _, expr := range t.expressions {
		if expressionStatement != "" {
			expressionStatement += ", "
		}
		s, v, err := expr.Build()
		if err != nil {
			return statement, values, err
		}
		expressionStatement += s
		expressionValues = append(expressionValues, v...)
	}
	statement += expressionStatement
	values = append(values, expressionValues...)
	if t.WhereWrapper != nil {
		where, whereValueHolder, err := t.WhereWrapper.Build()
		if err != nil {
			return "", nil, fmt.Errorf("syntax error: %v", err)
		}
		if where == "" {
			return statement, values, fmt.Errorf("missing where statement")
		}
		statement += " WHERE " + where
		values = append(values, whereValueHolder...)
	} else {
		return statement, values, fmt.Errorf("missing where statement")
	}
	values = placementValueFilter(values)
	return
}

func NewUpdateWrapper(model ...interface{}) UpdateWrapper {
	w := &updateWrapper{
		WhereWrapper: NewAndWhereWrapper(),
	}
	if len(model) > 0 {
		return w.Model(model[0])
	}
	return w
}
