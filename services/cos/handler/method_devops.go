package handler

import (
	"context"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *handler) Health(ctx context.Context, req *cos.HealthRequest) (resp *cos.HealthResponse, err error) {
	return &cos.HealthResponse{}, nil
}
