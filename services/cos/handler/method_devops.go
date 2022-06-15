package handler

import (
	"context"
	"github.com/juxuny/yc"
	cos "github.com/juxuny/yc/services/cos"
	"time"
)

func (t *handler) Health(ctx context.Context, req *cos.HealthRequest) (resp *cos.HealthResponse, err error) {
	return &cos.HealthResponse{
		CurrentTime: time.Now().Format(yc.DateTimeLayout),
	}, nil
}
