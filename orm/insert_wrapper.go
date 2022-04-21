package orm

import (
	"fmt"
	"github.com/juxuny/yc/utils"
	"reflect"
)

const DefaultInsertBatchSize = 1000

type InsertWrapper interface {
	Model(v interface{}) InsertWrapper
	BatchSize(batchSize int) InsertWrapper
	Fields(fields ...FieldName) InsertWrapper
	DataLen() int
	Reset() InsertWrapper
	Build() (statement string, values []interface{}, next bool, err error)
	Ignore() InsertWrapper
	Add(v ...interface{}) InsertWrapper
	OnDuplicatedUpdate(field FieldName, values ...interface{}) InsertWrapper
}

type insertWrapper struct {
	model             Model
	batchSize         int
	fields            []FieldName
	data              []interface{}
	iterator          int
	ignore            bool
	updateExpressions []UpdateExpression
}

func (t *insertWrapper) Add(v ...interface{}) InsertWrapper {
	t.data = append(t.data, v...)
	return t
}

func (t *insertWrapper) OnDuplicatedUpdate(field FieldName, values ...interface{}) InsertWrapper {
	if len(values) > 0 {
		t.updateExpressions = append(t.updateExpressions, NewSetValueUpdateExpression(field, values[0]))
	} else {
		t.updateExpressions = append(t.updateExpressions, NewDuplicatedUpdateExpression(field))
	}
	return t
}

func (t *insertWrapper) Ignore() InsertWrapper {
	t.ignore = true
	return t
}

func (t *insertWrapper) Build() (statement string, values []interface{}, next bool, err error) {
	if len(t.data) == 0 {
		return "", nil, false, fmt.Errorf("no data")
	}
	if t.ignore {
		statement += "INSERT IGNORE INTO"
	} else {
		statement += "INSERT INTO"
	}
	statement += " " + t.model.TableName.Wrap().String()
	fieldStatement := ""
	if len(t.fields) == 0 {
		t.fields = t.model.Fields
	}
	for _, f := range t.fields {
		if fieldStatement != "" {
			fieldStatement += ", "
		}
		fieldStatement += f.Wrap().String()
	}
	statement += " (" + fieldStatement + ") VALUES "
	for i := 0; t.iterator+i < len(t.data) && i < t.batchSize; i++ {
		valueMap := make(map[string]interface{})
		item := t.data[t.iterator+i]
		itemValue := reflect.ValueOf(item)
		if itemValue.Kind() == reflect.Ptr {
			itemValue = itemValue.Elem()
		}
		itemType := itemValue.Type()
		if itemValue.Kind() == reflect.Struct {
			for j := 0; j < itemValue.NumField(); j++ {
				f := itemValue.Field(j)
				ft := itemType.Field(j)
				fn := utils.ToUnderLine(ft.Name)
				if n, ok := ft.Tag.Lookup("orm"); ok {
					fn = n
				}
				if f.IsValid() {
					valueMap[wrap(fn, "`")] = f.Interface()
				}
			}
		} else {
			return "", nil, false, fmt.Errorf("unkonwn data type: %v", itemType.Name())
		}
		holderStatement := ""
		rowValues := make([]interface{}, 0)
		for _, f := range t.fields {
			if holderStatement != "" {
				holderStatement += ", "
			}
			if v, b := valueMap[f.String()]; b {
				holderStatement += "?"
				rowValues = append(rowValues, toPlacementValue(reflect.ValueOf(v)))
			} else {
				holderStatement += "NULL"
			}
		}
		statement += "(" + holderStatement + ")"
		values = append(values, rowValues...)
	}
	if len(t.updateExpressions) > 0 {
		updateStatement := ""
		updateValues := make([]interface{}, 0)
		for _, update := range t.updateExpressions {
			if updateStatement != "" {
				updateStatement += ", "
			}
			us, uv, err := update.Build()
			if err != nil {
				return "", nil, false, fmt.Errorf("update expression error: %v", err)
			}
			updateStatement += us
			updateValues = append(updateValues, uv...)
		}
		if updateStatement != "" {
			statement += " ON DUPLICATE KEY UPDATE " + updateStatement
			values = append(values, updateValues...)
		}
	}
	return
}

func (t *insertWrapper) Model(model interface{}) InsertWrapper {
	t.model = CreateModel(model)
	return t
}

func (t *insertWrapper) BatchSize(batchSize int) InsertWrapper {
	t.batchSize = batchSize
	return t
}

func (t *insertWrapper) Fields(fields ...FieldName) InsertWrapper {
	t.fields = fields
	return t
}

func (t *insertWrapper) DataLen() int {
	return len(t.data)
}

func (t *insertWrapper) Reset() InsertWrapper {
	t.iterator = 0
	return t
}

func NewInsertWrapper(model ...interface{}) InsertWrapper {
	w := &insertWrapper{
		batchSize: DefaultInsertBatchSize,
		iterator:  0,
	}
	if len(model) > 0 {
		return w.Model(model[0])
	}
	return w
}
