package handler

import (
	"context"
	logServer "github.com/juxuny/yc/services/log-server"
)

func (t *wrapper) Health(ctx context.Context, req *logServer.HealthRequest) (resp *logServer.HealthResponse, err error) {
	return t.handler.Health(ctx, req)
}
