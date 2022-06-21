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
	MdContextAccessKey     = "x-rpc-access-key"
	MdContextToken         = "x-rpc-token"
	MdContextSign          = "x-rpc-sign"
	MdContextSignMethod    = "x-rpc-sign-method"
	MdContextUserId        = "x-rpc-user-id"
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

func GetAccessKey(ctx context.Context) (accessKey string, found bool) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return "", found
	}
	v := md.Get(MdContextAccessKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func GetIncomingUserId(ctx context.Context) (userId *dt.ID, err error) {
	md := GetIncomingHeader(ctx)
	userIdString := md.Get(MdContextUserId)
	if len(userIdString) == 0 {
		return nil, errors.SystemError.RpcCallErrorNoUserId
	}
	uid, err := strconv.ParseUint(userIdString[0], 10, 64)
	if err != nil {
		return nil, errors.SystemError.RpcCallErrorNoUserId.Wrap(err)
	}
	return dt.NewIDPointer(uid), nil
}

func SetIncomingUserId(ctx context.Context, userId *dt.ID) (next context.Context) {
	if userId == nil || !userId.Valid {
		return ctx
	}
	md := GetIncomingHeader(ctx, metadata.New(map[string]string{
		MdContextUserId: userId.NumberAsString(),
	}))
	return metadata.NewIncomingContext(ctx, md)
}

func GetIncomingHeader(ctx context.Context, extra ...metadata.MD) metadata.MD {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return MergeMetadata(metadata.MD{}, extra...)
	}
	return MergeMetadata(md, extra...)
}

func GetOutgoingHeader(ctx context.Context, extra ...metadata.MD) metadata.MD {
	md, found := metadata.FromOutgoingContext(ctx)
	if !found {
		return MergeMetadata(metadata.MD{}, extra...)
	}
	return MergeMetadata(md, extra...)
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

func MergeMetadata(base metadata.MD, md ...metadata.MD) metadata.MD {
	ret := base.Copy()
	for _, m := range md {
		for k, v := range m {
			ret.Set(k, v...)
		}
	}
	return ret
}

func RpcCall(ctx context.Context, host string, queryPath string, data proto.Message, out interface{}, md metadata.MD, signHandler RpcSignContentHandler) (code int, err error) {
	body, err := proto.Marshal(data)
	if err != nil {
		return 0, errors.SystemError.RpcCallErrorIllegalRequestParams.Wrap(err)
	}
	output := bytes.NewBuffer(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, host+queryPath, output)
	if err != nil {
		return 0, errors.SystemError.RpcCallErrorNetwork.Wrap(err)
	}
	contextHeader := GetIncomingHeader(ctx)
	if signHandler != nil {
		method, sign, err := signHandler.Sum(body)
		if err != nil {
			return 0, err
		}
		md.Set(MdContextSign, sign)
		md.Set(MdContextSignMethod, string(method))
	}
	MergeRequestHeaderFromMetadata(req, contextHeader, md)
	req.Header.Set("content-type", "application/protobuf")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error(err)
		return 0, errors.SystemError.RpcCallErrorBuildRequest.Wrap(err)
	}
	code = resp.StatusCode
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, errors.SystemError.RpcCallErrorReadResponseBody.Wrap(err)
	}
	ct := resp.Header.Get("content-type")
	if ct == "" {
		ct = resp.Header.Get("Content-Type")
	}
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
