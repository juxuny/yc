package handler

import (
	"context"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *handler) SaveConfig(ctx context.Context, req *cos.SaveConfigRequest) (resp *cos.SaveConfigResponse, err error) {
	return &cos.SaveConfigResponse{}, nil
}

func (t *handler) DeleteConfig(ctx context.Context, req *cos.DeleteConfigRequest) (resp *cos.DeleteConfigResponse, err error) {
	return &cos.DeleteConfigResponse{}, nil
}

func (t *handler) ListConfig(ctx context.Context, req *cos.ListConfigRequest) (resp *cos.ListConfigResponse, err error) {
	return &cos.ListConfigResponse{}, nil
}
