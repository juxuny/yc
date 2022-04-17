package trace

import (
	"context"
	"fmt"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/env"
	"github.com/petermattis/goid"
	"google.golang.org/grpc/metadata"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
)

var (
	metadataMap = &sync.Map{}
	waitGroup   = &sync.WaitGroup{} // 系统退出的时候等所有协程结束之后才退出
)

const (
	HeaderTraceReqId = "trace-req-id"
	HeaderTraceUid   = "trace-uid"
)

func init() {
	md := metadata.New(map[string]string{
		HeaderTraceReqId: "main",
	})
	InitContext(md)
}

func GetMetadata(goId ...int64) metadata.MD {
	var id int64
	if len(goId) > 0 {
		id = goId[0]
	} else {
		id = goid.Get()
	}
	if v, ok := metadataMap.Load(id); ok {
		return v.(metadata.MD)
	}
	return metadata.New(map[string]string{})
}

func InitReqId(reqId ...string) {
	if len(reqId) > 0 && reqId[0] != "" {
		md := metadata.New(map[string]string{
			HeaderTraceReqId: reqId[0],
		})
		InitContext(md)
	} else {
		InitContext()
	}
}

func InitContext(initMetadata ...metadata.MD) {
	id := goid.Get()
	md := make(metadata.MD)
	for _, item := range initMetadata {
		for k, v := range item {
			md[k] = v
		}
	}
	reqId := md[HeaderTraceReqId]
	if len(reqId) == 0 {
		reqId := env.DefaultEnv.Mode + "." + genReqId(20)
		md[HeaderTraceReqId] = []string{reqId}
	}
	metadataMap.Store(id, md)
	return
}

func Clean() {
	id := goid.Get()
	metadataMap.Delete(id)
}

func SetUid(uid dt.ID) {
	id := goid.Get()
	md := GetMetadata(id)
	md.Set(HeaderTraceUid, fmt.Sprintf("%d", uid.Int64))
	metadataMap.Store(id, md)
}

func GetUid() (ret dt.ID) {
	md := GetMetadata()
	v := md.Get(HeaderTraceUid)
	if len(v) > 0 {
		id, err := strconv.ParseInt(v[0], 10, 64)
		if err != nil {
			return
		}
		ret = dt.NewID(id)
	}
	return
}

func GetReqId(inputId ...int64) string {
	md := GetMetadata(inputId...)
	ret := md.Get(HeaderTraceReqId)
	if len(ret) > 0 {
		return ret[0]
	}
	return ""
}

func GoRun(f func()) {
	id := goid.Get()
	parentMetadata := GetMetadata(id)
	waitGroup.Add(1)
	go func(parentId int64, parentMetadata metadata.MD) {
		defer func() {
			waitGroup.Done()
		}()
		reqId := genReqId(6)
		parentReqId := parentMetadata.Get(HeaderTraceReqId)
		childMetadata := parentMetadata.Copy()
		newReqId := fmt.Sprintf("%s.%s", strings.Join(parentReqId, ""), reqId)
		childMetadata.Set(HeaderTraceReqId, newReqId)
		InitContext(childMetadata)
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
			}
		}()
		f()
	}(id, parentMetadata)
}

func Wait() {
	waitGroup.Wait()
}

func WithContext(ctx context.Context) {
	md, b := metadata.FromIncomingContext(ctx)
	if b {
		InitContext(md)
	} else {
		InitContext()
	}
}

func GoRunWithContext(ctx context.Context, f func(ctx context.Context)) {
	GoRun(func() {
		f(ctx)
	})
}
