package yc

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	"google.golang.org/grpc/metadata"
	"strconv"
)

const (
	MdContextCallerService = "x-rpc-caller-service"
	MdContextCallerLevel   = "x-rpc-caller-level"
	MdContextToken         = "x-rpc-token"
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

func GetTokenFromContext(ctx context.Context) (token string, err error) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return "", errors.SystemError.NotFoundRpcToken
	}
	v := md.Get(MdContextToken)
	if len(v) == 0 {
		return "", errors.SystemError.NotFoundRpcToken
	}
	return v[0], nil
}

func ParseJwt(ctx context.Context) (*jwt.Claims, error) {
	token, err := GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return jwt.ParseToken(token)
}

func GetUserId(ctx context.Context) (userId dt.ID, err error) {
	claims, err := ParseJwt(ctx)
	if err != nil {
		log.Error(err)
		return dt.InvalidID(), err
	}
	return claims.UserId, nil
}
