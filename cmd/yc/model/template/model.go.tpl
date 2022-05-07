package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/orm"
	{{.PackageAlias}} "{{.GoModuleName}}"
)

var {{.TableName}} = {{.TableName|lowerFirst}} {
{{range $field := .Fields}}	{{.FieldName|upperFirst}}: orm.FieldName("{{.OrmFieldName}}"),
{{end}}}

type {{.ModelName}} struct {
{{range $field := .Fields}}	{{.FieldName|upperFirst}} {{.ModelDataType}} `json:"{{.FieldName|lowerFirst}}" orm:"{{.OrmFieldName}}"`
{{end}}}

func ({{.ModelName}}) TableName() string {
	return cos.Name + "_" + "{{.TableNameWithoutServicePrefix}}"
}

{{$modelName := .ModelName}}
{{$packageAlias := .PackageAlias}}
{{range $ref := .Refs}}
func (t {{$modelName}}) To{{$ref.ModelName}}() {{$packageAlias}}.{{$ref.ModelName}} {
	return {{$packageAlias}}.{{$ref.ModelName}}{
	{{range $refField := $ref.Fields}}	{{$refField.FieldName|upperFirst}}: t.{{$refField.FieldName|upperFirst}},
	{{end}}}
}{{end}}

type {{.TableName|lowerFirst}} struct {
{{range $field := .Fields}}	{{.FieldName|upperFirst}} orm.FieldName
{{end}}}

func ({{.TableName|lowerFirst}}) TableName() string {
	return cos.Name + "_" + "{{.TableNameWithoutServicePrefix}}"
}
{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) FindOneBy{{$field.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (data {{$field.ModelName}}, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq({{$field.TableName}}.{{$field.FieldName|upperFirst}}, {{$field.FieldName|lowerFirst}}){{if $field.HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{$field.TableName}}.DeletedAt, 0).IsNull({{$field.TableName}}.DeletedAt)){{end}}
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}
{{end}}{{end}}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) UpdateBy{{$field.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper({{.ModelName}}{})
	w.Eq({{.TableName}}.{{$field.FieldName|upperFirst}}, {{$field.FieldName|lowerFirst}}){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
{{end}}{{end}}

func ({{.TableName|lowerFirst}}) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper({{.ModelName}}{})
	w.SetWhere(where).Updates(update){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) DeleteBy{{$field.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper({{.ModelName}}{})
	w.Eq({{.TableName}}.{{$field.FieldName|upperFirst}}, {{$field.FieldName|lowerFirst}})
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
{{end}}{{end}}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) SoftDeleteBy{{$field.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper({{.ModelName}}{})
	w.SetValue({{.TableName}}.DeletedAt, orm.Now())
	w.Eq({{.TableName}}.{{$field.FieldName|upperFirst}}, {{$field.FieldName|lowerFirst}})
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
{{end}}{{end}}

func ({{.TableName|lowerFirst}}) Find(ctx context.Context, where orm.WhereWrapper) (list []{{.ModelName}}, err error) {
	w := orm.NewQueryWrapper({{.ModelName}}{})
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) FindBy{{.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (list []{{.ModelName}}, err error) {
	w := orm.NewQueryWrapper({{.ModelName}}{})
	w.Eq({{.TableName}}.{{.FieldName|upperFirst}}, {{.FieldName|lowerFirst}}){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}
{{end}}{{end}}

func ({{.TableName|lowerFirst}}) Page(ctx context.Context, pageNum, pageSize int, where orm.WhereWrapper) (list []{{.ModelName}}, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) PageBy{{.FieldName|upperFirst}}(ctx context.Context, pageNum, pageSize int, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (list []{{.ModelName}}, err error) {
	w := orm.NewQueryWrapper({{.ModelName}}{})
	w.Eq({{.TableName}}.{{.FieldName|upperFirst}}, {{.FieldName|lowerFirst}}){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}
{{end}}{{end}}

func ({{.TableName|lowerFirst}}) Count(ctx context.Context, where orm.WhereWrapper) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

{{range $field := .Fields}}{{if $field.HasIndex}}
func ({{.TableName|lowerFirst}}) CountBy{{.FieldName|upperFirst}}(ctx context.Context, {{$field.FieldName|lowerFirst}} {{$field.ModelDataType}}) (count int, err error) {
	w := orm.NewQueryWrapper({{.ModelName}}{})
	w.Eq({{.TableName}}.{{.FieldName|upperFirst}}, {{.FieldName|lowerFirst}}){{if .HasDeletedAt}}
	w.Nested(orm.NewOrWhereWrapper().Eq({{.TableName}}.DeletedAt, 0).IsNull({{.TableName}}.DeletedAt)){{end}}
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}
{{end}}{{end}}

func ({{.TableName|lowerFirst}}) Create(ctx context.Context, data ...{{.ModelName}}) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper({{.ModelName}}{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
