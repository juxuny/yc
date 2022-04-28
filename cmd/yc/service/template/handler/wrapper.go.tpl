package handler

import {{.PackageAlias}} "{{.GoModuleName}}"

type wrapper struct {
	handler *handler
	{{.PackageAlias}}.Unimplemented{{.ServiceStruct}}Server
}

func NewWrapper() *wrapper {
	return &wrapper{
		handler: &handler{},
	}
}
