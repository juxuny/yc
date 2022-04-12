package handler

import (
	"context"
	logServer "github.com/juxuny/yc/services/log-server"
)

func (t *wrapper) Print(ctx context.Context, req *logServer.PrintRequest) (resp *logServer.PrintResponse, err error) {
	return t.handler.Print(ctx, req)
}
