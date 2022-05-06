package services

import (
	"github.com/juxuny/yc/utils"
	"html/template"
	"strings"
)

type baseEntity struct{}

func (baseEntity) Lt() template.HTML {
	return "<"
}

func (baseEntity) Gt() template.HTML {
	return ">"
}

func (baseEntity) Le() template.HTML {
	return "<="
}

func (baseEntity) Ge() template.HTML {
	return ">="
}

type ServiceEntity struct {
	baseEntity
	ServiceName   string
	PackageName   string
	ServiceStruct string
	ServiceDir    string
	ProtoFileName string
	PackageAlias  string
}

func NewServiceEntity(serviceName string) ServiceEntity {
	ret := ServiceEntity{
		ServiceDir:    strings.ReplaceAll(utils.ToUnderLine(serviceName), "_", "-"),
		ServiceName:   serviceName,
		PackageName:   utils.ToUnderLine(serviceName),
		ProtoFileName: utils.ToUnderLine(serviceName),
		ServiceStruct: utils.ToHump(serviceName),
		PackageAlias:  utils.ToLowerFirst(utils.ToHump(serviceName)),
	}
	return ret
}

type ValidatorEntity struct {
	ServiceEntity
	Messages []Message
}

type ValidatorFormula struct {
	Pattern  string
	RefValue string
}

type MessageField struct {
	Name      string
	Formulas  []ValidatorFormula
	ParamName string
	Error     string
}

type Message struct {
	Name   string
	Fields []MessageField
}

type HandlerInitEntity struct {
	baseEntity
	ServiceEntity
	GoModuleName string
}

type MethodEntity struct {
	HandlerInitEntity
	MethodName string
	Request    string
	Response   string
	UseAuth    bool
}

type EntrypointEntity struct {
	baseEntity
	ServiceEntity
	GoModuleName string
}

type ModelField struct {
	FieldName    string
	OrmFieldName string
	HasIndex     bool
}

type Model struct {
	Fields       []ModelField
	TableName    string
	HasDeletedAt bool
}

type ModelEntity struct {
	ServiceEntity
	NoDuplicatedFields []ModelField
	Models             []Model
}
