package orm

import (
	"fmt"
	"github.com/juxuny/yc/utils"
	"path"
	"reflect"
	"strings"
)

type JoinStatement string

const (
	JoinStatementLeftJoin  = JoinStatement("LEFT JOIN")
	JoinStatementRightJoin = JoinStatement("RIGHT JOIN")
)

type Join struct {
	Statement    JoinStatement
	TableName    TableName
	WhereWrapper WhereWrapper
}

type QueryWrapper interface {
	Model(v interface{}) QueryWrapper
	WhereWrapper
	OrderDesc(fields ...FieldName) QueryWrapper
	OrderAsc(fields ...FieldName) QueryWrapper
	Offset(offset int) QueryWrapper
	Limit(limit int) QueryWrapper
	Page(pageNum, pageSize int) QueryWrapper
	Select(fields ...FieldName) QueryWrapper
	LeftJoin(tableName TableName, where WhereWrapper) QueryWrapper
	RightJoin(tableName TableName, where WhereWrapper) QueryWrapper
}

type queryWrapper struct {
	WhereWrapper
	model             Model
	offset, limit     int
	orderStatement    string
	selectFields      []FieldName
	selectValueHolder []interface{}
	joinList          []Join
}

func (t *queryWrapper) LeftJoin(tableName TableName, where WhereWrapper) QueryWrapper {
	t.joinList = append(t.joinList, Join{
		Statement:    JoinStatementLeftJoin,
		TableName:    tableName,
		WhereWrapper: where,
	})
	return t
}

func (t *queryWrapper) RightJoin(tableName TableName, where WhereWrapper) QueryWrapper {
	t.joinList = append(t.joinList, Join{
		Statement:    JoinStatementRightJoin,
		TableName:    tableName,
		WhereWrapper: where,
	})
	return t
}

func (t *queryWrapper) Select(fields ...FieldName) QueryWrapper {
	t.selectFields = fields
	return t
}

func (t *queryWrapper) SelectWithHolder(fields []FieldName, v ...interface{}) QueryWrapper {
	t.selectFields = fields
	t.selectValueHolder = v
	return t
}

func (t *queryWrapper) OrderDesc(fields ...FieldName) QueryWrapper {
	if t.orderStatement != "" {
		t.orderStatement += " AND "
	}
	t.orderStatement += strings.TrimSpace(JoinFieldNames(fields, ", ")) + " DESC"
	return t
}

func (t *queryWrapper) OrderAsc(fields ...FieldName) QueryWrapper {
	if t.orderStatement != "" {
		t.orderStatement += " AND "
	}
	t.orderStatement += strings.TrimSpace(JoinFieldNames(fields, ", ")) + " ASC"
	return t
}

func (t *queryWrapper) Offset(offset int) QueryWrapper {
	t.offset = offset
	return t
}

func (t *queryWrapper) Limit(limit int) QueryWrapper {
	t.limit = limit
	return t
}

func (t *queryWrapper) Page(pageNum, pageSize int) QueryWrapper {
	t.offset = (pageNum - 1) * pageSize
	t.limit = pageSize
	return t
}

func (t *queryWrapper) Order(asc bool, fields ...FieldName) QueryWrapper {
	if asc {
		return t.OrderAsc(fields...)
	} else {
		return t.OrderDesc(fields...)
	}
}

func (t *queryWrapper) Model(v interface{}) QueryWrapper {
	modelValue := reflect.ValueOf(v)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	m := modelValue.MethodByName("TableName")
	if m.IsValid() {
		out := m.Call(nil)
		for _, v := range out {
			if v.Kind() == reflect.String {
				t.model.TableName = TableName(TablePrefix + "_" + v.String())
			}
		}
	}
	if t.model.TableName == "" {
		n := modelValue.Type().Name()
		if strings.Contains(n, ".") {
			n = path.Ext(n)
		}
		tn := utils.ToUnderLine(n)
		t.model.TableName = TableName(tn).Prefix(TablePrefix)
	}
	modelType := modelValue.Type()
	for i := 0; i < modelType.NumField(); i++ {
		f := modelType.Field(i)
		tag := f.Tag
		fieldName := ""
		if fn, ok := tag.Lookup("orm"); ok {
			fieldName = strings.TrimSpace(fn)
		} else {
			fieldName = utils.ToUnderLine(f.Name)
		}
		t.model.Fields = append(t.model.Fields, FieldName(fieldName).Wrap())
	}
	return t
}

func (t *queryWrapper) Build() (statement string, values []interface{}, err error) {
	statement = "SELECT "
	// select fields
	selectStatement := ""
	if len(t.selectFields) == 0 {
		selectStatement = strings.TrimSpace(JoinFieldNames(t.model.Fields, ", "))
	} else {
		selectStatement = strings.TrimSpace(JoinFieldNames(t.selectFields, ", "))
		values = append(values, t.selectValueHolder...)
	}
	statement += selectStatement
	// from
	statement += " FROM " + t.model.TableName.Wrap().String()
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
	if t.WhereWrapper != nil {
		where, whereValueHolder, err := t.WhereWrapper.Build()
		if err != nil {
			return "", nil, fmt.Errorf("syntax error: %v", err)
		}
		statement += " WHERE " + where
		values = append(values, whereValueHolder...)
	}
	statement += t.orderStatement
	if t.limit > 0 {
		statement += " LIMIT ? OFFSET ?"
		values = append(values, t.limit, t.offset)
	}
	return
}

func NewQueryWrapper(model ...interface{}) QueryWrapper {
	w := &queryWrapper{
		WhereWrapper: NewAndWhereWrapper(),
	}
	if len(model) > 0 {
		return w.Model(model[0])
	}
	return w
}
