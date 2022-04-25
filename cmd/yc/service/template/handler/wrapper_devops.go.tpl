package handler

import (
	"context"
	logServer "github.com/juxuny/yc/services/log-server"
	"github.com/juxuny/yc/trace"
)

func (t *wrapper) Health(ctx context.Context, req *logServer.HealthRequest) (resp *logServer.HealthResponse, err error) {
	trace.WithContext(ctx)
	defer trace.Clean()
	return t.handler.Health(ctx, req)
}
