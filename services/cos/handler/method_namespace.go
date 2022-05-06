package handler

import (
	"context"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *handler) SaveNamespace(ctx context.Context, req *cos.SaveNamespaceRequest) (resp *cos.SaveNamespaceResponse, err error) {
	return &cos.SaveNamespaceResponse{}, nil
}

func (t *handler) ListNamespace(ctx context.Context, req *cos.ListNamespaceRequest) (resp *cos.ListNamespaceResponse, err error) {
	return &cos.ListNamespaceResponse{}, nil
}

func (t *handler) DeleteNamespace(ctx context.Context, req *cos.DeleteNamespaceRequest) (resp *cos.DeleteNamespaceResponse, err error) {
	return &cos.DeleteNamespaceResponse{}, nil
}
