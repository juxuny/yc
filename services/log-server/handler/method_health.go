package handler

import (
	"context"
	"github.com/juxuny/yc"
	logServer "github.com/juxuny/yc/services/log-server"
	"time"
)

func (t *Handler) Health(ctx context.Context, req *logServer.HealthRequest) (resp *logServer.HealthResponse, err error) {
	return &logServer.HealthResponse{CurrentTime: time.Now().Format(yc.DateTimeLayout)}, nil
}
