package middle

import (
	"context"
	"github.com/juxuny/yc/errors"
	"runtime/debug"
)

type RecoverHandler struct{}

func (t *RecoverHandler) Run(context context.Context) (isEnd bool, err error) {
	if err := recover(); err != nil {
		debug.PrintStack()
		return true, errors.SystemError.InternalError.WithField("panic", err)
	}
	return
}
