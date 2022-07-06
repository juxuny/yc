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
	Version       string
}

func NewServiceEntity(serviceName string, version string) ServiceEntity {
	ret := ServiceEntity{
		ServiceDir:    strings.ReplaceAll(utils.ToUnderLine(serviceName), "_", "-"),
		ServiceName:   serviceName,
		PackageName:   utils.ToUnderLine(serviceName),
		ProtoFileName: utils.ToUnderLine(serviceName),
		ServiceStruct: utils.ToHump(serviceName),
		PackageAlias:  utils.ToLowerFirst(utils.ToHump(serviceName)),
		Version:       version,
	}
	return ret
}

type ValidatorEntity struct {
	ServiceEntity
	Messages []ValidatorMessage
}

type ValidatorFormula struct {
	Pattern  string
	RefValue string
}

type ValidatorMessageField struct {
	Name      string
	Formulas  []ValidatorFormula
	ParamName string
	Error     string
}

type ValidatorMessage struct {
	Name   string
	Fields []ValidatorMessageField
	Desc   string
}

type HandlerInitEntity struct {
	baseEntity
	ServiceEntity
	GoModuleName string
}

type MethodEntity struct {
	HandlerInitEntity
	Group      string
	MethodName string
	Request    string
	Response   string
	UseAuth    bool
	Desc       string
	Api        string
}

type EntrypointEntity struct {
	baseEntity
	ServiceEntity
	GoModuleName string
}

type ModelField struct {
	TableName      string
	ModelName      string
	FieldName      string
	OrmFieldName   string
	JsonFieldName  string
	ModelDataType  string
	Ignore         bool
	HasIndex       bool
	HasUnique      bool
	HasDeletedAt   bool
	CSharpDataType string
}

type Model struct {
	Fields                        []ModelField
	TableName                     string
	ModelName                     string
	TableNameWithoutServicePrefix string
	HasDeletedAt                  bool
	Refs                          []RefModel
}

func (t Model) ToRefModel() RefModel {
	return RefModel{
		ModelName: t.ModelName,
		Fields:    t.Fields,
	}
}

type RefModel struct {
	ModelName string
	Fields    []ModelField
}

type ModelEntity struct {
	ServiceEntity
	GoModuleName         string
	CSharpModelNamespace string
	CSharpBaseNamespace  string
	Model
}

type ClientSdkEntity struct {
	Methods []MethodEntity
	ServiceEntity
}

type CloneMessage struct {
	Name string
}

type CloneEntity struct {
	ServiceEntity
	Messages []CloneMessage
}
