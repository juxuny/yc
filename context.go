package yc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
		return 0, errors.SystemError.RpcCallErrorInvalidCallerLevel.Wrap(err)
	}
	return int(levelInt64), nil
}

func GetCallerLevelFromContext(ctx context.Context) (int, error) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return 0, errors.SystemError.RpcCallErrorMetaEmpty
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
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return nil, errors.SystemError.InvalidToken.Wrap(err)
	}
	return claim, nil
}

func GetUserId(ctx context.Context) (userId dt.ID, err error) {
	claims, err := ParseJwt(ctx)
	if err != nil {
		log.Error(err)
		return dt.InvalidID(), err
	}
	return claims.UserId, nil
}

func GetHeader(ctx context.Context) metadata.MD {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return make(metadata.MD)
	}
	return md
}

func MergeRequestHeaderFromMetadata(req *http.Request, md ...metadata.MD) {
	for _, m := range md {
		for k, values := range m {
			for i, v := range values {
				if i == 0 {
					req.Header.Set(k, v)
				} else {
					req.Header.Add(k, v)
				}
			}
		}
	}
}

func RpcCall(ctx context.Context, host string, queryPath string, data proto.Message, out interface{}, md metadata.MD) (code int, err error) {
	body, err := proto.Marshal(data)
	if err != nil {
		return 0, errors.SystemError.RpcCallErrorIllegalRequestParams.Wrap(err)
	}
	output := bytes.NewBuffer(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, host+queryPath, output)
	if err != nil {
		return 0, errors.SystemError.RpcCallErrorNetwork.Wrap(err)
	}
	contextHeader := GetHeader(ctx)
	MergeRequestHeaderFromMetadata(req, contextHeader, md)
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, errors.SystemError.RpcCallErrorBuildRequest.Wrap(err)
	}
	code = resp.StatusCode
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, errors.SystemError.RpcCallErrorReadResponseBody.Wrap(err)
	}
	ct := resp.Header.Get("content-type")
	if strings.Contains(ct, "json") {
		if code == http.StatusOK {
			err = json.Unmarshal(respData, out)
		} else {
			var respError errors.Error
			err = json.Unmarshal(respData, &respError)
			if err != nil {
				return code, err
			} else {
				return code, respError
			}
		}
	} else if strings.Contains(ct, "protobuf") {
		if message, ok := out.(proto.Message); ok {
			if code == http.StatusOK {
				err = proto.Unmarshal(respData, message)
			} else {
				var respError = &dt.Error{}
				err = proto.Unmarshal(respData, respError)
				if err != nil {
					return code, err
				} else {
					return code, respError.Error()
				}
			}
		} else {
			return code, errors.SystemError.InvalidProtobufHolder
		}
	} else {
		return code, errors.SystemError.InvalidContentType.Wrap(fmt.Errorf("Content-Type: %s", ct))
	}
	return
}
