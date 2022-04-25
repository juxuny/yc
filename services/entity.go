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
