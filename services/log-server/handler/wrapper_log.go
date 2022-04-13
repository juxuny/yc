package handler

import (
	"context"
	"github.com/juxuny/yc/log"
	logServer "github.com/juxuny/yc/services/log-server"
	"github.com/juxuny/yc/trace"
)

func (t *wrapper) Print(ctx context.Context, req *logServer.PrintRequest) (resp *logServer.PrintResponse, err error) {
	trace.WithContext(ctx)
	defer trace.Clean()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	return t.handler.Print(ctx, req)
}
