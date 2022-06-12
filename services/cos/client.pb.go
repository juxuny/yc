package cos

import (
	"context"
	"github.com/juxuny/yc"
	"google.golang.org/grpc/metadata"
)

type Client interface {
}

type client struct{}

func (*client) Health(ctx context.Context, req *HealthRequest) (resp *HealthResponse, err error) {
	md := yc.GetHeader(ctx)
	metadata.NewOutgoingContext(ctx, md)
	return
}
