package cos

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/errors"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	"time"
)

type Client interface {
	Health(ctx context.Context, req *HealthRequest) (resp *HealthResponse, err error)
}

type client struct {
	Service      string
	Entrypoint   []string
	randInstance *rand.Rand
}

var DefaultClient Client

func NewClient(entrypoint ...string) Client {
	return &client{
		Service:      Name,
		Entrypoint:   entrypoint,
		randInstance: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (t *client) selectEntrypoint() (entrypoint string) {
	t.randInstance.Seed(time.Now().UnixNano())
	if len(t.Entrypoint) == 0 {
		return ""
	}
	return t.Entrypoint[t.randInstance.Intn(len(t.Entrypoint))]
}

func (t *client) Health(ctx context.Context, req *HealthRequest) (resp *HealthResponse, err error) {
	md := yc.GetHeader(ctx)
	metadata.NewOutgoingContext(ctx, md)
	resp = &HealthResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.selectEntrypoint(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}
