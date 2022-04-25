package handler

import logServer "github.com/juxuny/yc/services/log-server"

type wrapper struct {
	handler *handler
	logServer.UnimplementedLogServerServer
}

func NewWrapper() *wrapper {
	return &wrapper{
		handler: &handler{},
	}
}
