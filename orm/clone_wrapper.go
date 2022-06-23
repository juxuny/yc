package orm

import "fmt"

type CloneWrapper interface {
	Build() (statement string, values []interface{}, err error)
}

type cloneWrapper struct {
	templateTableName TableName
	newTableName      TableName
}

func (t *cloneWrapper) Build() (statement string, values []interface{}, err error) {
	return fmt.Sprintf("CREATE TABLE %s LIKE %s", t.newTableName.Wrap().String(), t.templateTableName.Wrap().String()), nil, nil
}

func NewCloneWrapper(templateTableName, newTableName TableName) CloneWrapper {
	return &cloneWrapper{
		templateTableName: templateTableName,
		newTableName:      newTableName,
	}
}
