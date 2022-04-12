package handler

import (
	"context"
	logServer "github.com/juxuny/yc/services/log-server"
)

func (t *wrapper) Print(ctx context.Context, req *logServer.PrintRequest) (resp *logServer.PrintResponse, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return t.handler.Print(ctx, req)
}
