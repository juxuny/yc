package services

import (
	"github.com/juxuny/yc/utils"
	"strings"
)

type ServiceEntity struct {
	ServiceName   string
	PackageName   string
	ServiceStruct string
	ServiceDir    string
	ProtoFileName string
}

func NewServiceEntity(serviceName string) ServiceEntity {
	ret := ServiceEntity{
		ServiceDir:    strings.ReplaceAll(utils.ToUnderLine(serviceName), "_", "-"),
		ServiceName:   serviceName,
		PackageName:   utils.ToUnderLine(serviceName),
		ProtoFileName: utils.ToUnderLine(serviceName),
		ServiceStruct: utils.ToHump(serviceName),
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
