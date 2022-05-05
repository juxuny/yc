package yc

import (
	"context"
	"github.com/juxuny/yc/errors"
	"google.golang.org/grpc/metadata"
	"strconv"
)

const (
	MdContextCallerService = "x-rpc-caller-service"
	MdContextCallerLevel   = "x-rpc-caller-level"
)

func GetCallerLevelFromMd(md metadata.MD) (level int, err error) {
	v := md.Get(MdContextCallerLevel)
	if len(v) == 0 {
		return 0, errors.SystemError.NotFoundRpcCallerLevel
	}
	levelInt64, err := strconv.ParseInt(v[0], 10, 64)
	if err != nil {
		return 0, errors.SystemError.InvalidRpcCallerLevel.Wrap(err)
	}
	return int(levelInt64), nil
}

func GetCallerLevelFromContext(ctx context.Context) (int, error) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return 0, errors.SystemError.RpcCallMetaEmpty
	}
	return GetCallerLevelFromMd(md)
}
