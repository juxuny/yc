package handler

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/services/cos"
	"time"
)

func (h *Handler) Health(ctx context.Context, req *cos.HealthRequest) (resp *cos.HealthResponse, err error) {
	return &cos.HealthResponse{CurrentTime: time.Now().Format(yc.DateTimeLayout)}, nil
}
