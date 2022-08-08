package middle

import (
	"context"
	"github.com/juxuny/yc/router"
)

type Handler interface {
	Run(context context.Context) (nextContext context.Context, isEnd bool, err error)
}

type Group interface {
	Run(context context.Context) (nextContext context.Context, isEnd bool, err error)
	Add(h ...Handler) Group
}

type group struct {
	handlers []Handler
}

func (t *group) Run(context context.Context) (next context.Context, isEnd bool, err error) {
	for _, h := range t.handlers {
		next, isEnd, err = h.Run(context)
		if next != nil {
			context = next
		}
		if err != nil {
			return context, true, err
		}
		if isEnd {
			return context, isEnd, nil
		}
	}
	return context, false, nil
}

func (t *group) Add(h ...Handler) Group {
	t.handlers = append(t.handlers, h...)
	return t
}

func NewGroup(h ...Handler) Group {
	ret := &group{
		handlers: h,
	}
	return ret
}

type apiHandler struct {
	f func(ctx context.Context)
}

func (t *apiHandler) Run(ctx context.Context) (next context.Context, isEnd bool, err error) {
	t.f(ctx)
	return ctx, false, nil
}

func NewApiHandler(f func(ctx context.Context)) Handler {
	return &apiHandler{f: f}
}

type validatorHandler struct {
	router.ValidatorHandler
}

func (t *validatorHandler) Run(ctx context.Context) (next context.Context, isEnd bool, err error) {
	err = t.Validate()
	return ctx, err != nil, err
}

func NewValidatorHandler(v router.ValidatorHandler) Handler {
	return &validatorHandler{
		ValidatorHandler: v,
	}
}
