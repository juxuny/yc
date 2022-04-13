package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	logServer "github.com/juxuny/yc/services/log-server"
	"github.com/juxuny/yc/services/log-server/config"
	"github.com/juxuny/yc/services/log-server/impl"
	"strings"
)

var fileLogger impl.Logger

func Flush() error {
	if fileLogger != nil {
		return fileLogger.Flush()
	}
	return nil
}

func (t *handler) Print(ctx context.Context, req *logServer.PrintRequest) (resp *logServer.PrintResponse, err error) {
	if config.Env.LogDir == "" {
		return nil, errors.SystemError.LogDirEmpty
	}
	if fileLogger == nil {
		fileLogger = impl.NewDefaultFileLogger(config.Env.LogDir, config.Env.CacheSize, config.Env.FlushSeconds)
	}
	resp = &logServer.PrintResponse{}
	data := "[" + req.GetDateTime() + "] " + fmt.Sprintf("<req_id=%s>", req.GetReqId()) + " " + "[" + req.GetLevel().Name() + "] " + "(" + req.GetFileLine() + ") "
	if len(req.Extra) > 0 {
		data += "["
		for _, item := range req.GetExtra() {
			data += fmt.Sprintf("%s:%s ", item.Name, item.Value)
		}
		data = strings.TrimSpace(data)
		data += "] "
	}
	data += "message: " + req.GetContent()
	log.Debug(req.GetContent())
	err = fileLogger.Info(config.Env.FilePrefix, data)
	return
}
