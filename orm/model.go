package orm

import "strings"

var TablePrefix = ""

type Model struct {
	TableName TableName
	Fields    []FieldName
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
