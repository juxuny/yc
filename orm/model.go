package orm

import (
	"github.com/juxuny/yc/utils"
	"path"
	"reflect"
	"strings"
)

var TablePrefix = ""

type Model struct {
	TableName TableName
	Fields    []FieldName
}

func CreateModel(model interface{}) Model {
	ret := Model{}
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	m := modelValue.MethodByName("TableName")
	if m.IsValid() {
		out := m.Call(nil)
		for _, v := range out {
			if v.Kind() == reflect.String {
				ret.TableName = TableName(strings.Trim(TablePrefix+"_"+v.String(), "_"))
			}
		}
	}
	if ret.TableName == "" {
		n := modelValue.Type().Name()
		if strings.Contains(n, ".") {
			n = path.Ext(n)
		}
		tn := utils.ToUnderLine(n)
		ret.TableName = TableName(tn).Prefix(TablePrefix)
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
		ret.Fields = append(ret.Fields, FieldName(fieldName).Wrap())
	}
	return ret
}

type TableName string

func (t TableName) Prefix(p string) TableName {
	s := strings.Trim(p+"_"+string(t), "_")
	return TableName(strings.ReplaceAll(s, "__", "_"))
}

func (t TableName) Suffix(p string) TableName {
	s := strings.Trim(string(t)+"_"+p, "_")
	return TableName(strings.ReplaceAll(s, "__", "_"))
}

func (t TableName) Wrap() TableName {
	s := string(t)
	if strings.Contains(s, ".") {
		return t
	}
	if strings.Contains(s, "`") {
		return t
	}
	return TableName(wrap(s, "`"))
}

func (t TableName) String() string {
	return string(t)
}

func (t TableName) Alias(name string) string {
	return t.String() + " " + name
}

type FieldName string

func (t FieldName) String() string {
	return string(t)
}

func (t FieldName) WithTableAlias(alias string) FieldName {
	return FieldName(strings.Trim(alias, ".")) + t
}

func (t FieldName) Wrap() FieldName {
	s := string(t)
	if strings.Contains(s, ".") {
		return t
	}
	if strings.Contains(s, "`") {
		return t
	}
	return FieldName(wrap(s, "`"))
}

func (t FieldName) WithAlias(alias string) FieldName {
	return t + " AS " + FieldName(alias)
}
