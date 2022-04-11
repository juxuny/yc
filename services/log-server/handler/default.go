package handler

import (
	logServer "github.com/juxuny/yc/services/log-server"
)

type Handler struct {
	logServer.UnimplementedLogServerServer
}
