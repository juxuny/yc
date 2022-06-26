package router

import (
	"net/http"
)

type Router struct {
	prefix             string
	handlerChain       []HandleFunc
	mainHandlerWrapper *mainHandlerWrapper
}

func NewRouter(prefix string, handlerChain ...HandleFunc) *Router {
	hc := make([]HandleFunc, 0)
	hc = append(hc, handlerChain...)
	return &Router{
		prefix:             prefix,
		handlerChain:       hc,
		mainHandlerWrapper: newMainHandlerWrapper(prefix),
	}
}

func (t *Router) Register(groupName string, handler interface{}) error {
	return t.mainHandlerWrapper.register(groupName, handler)
}

func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(t.handlerChain, w, r)
	ctx.Next()
}

func (t *Router) Start(addr string) error {
	t.handlerChain = append(t.handlerChain, t.mainHandlerWrapper.callerHandleFunc)
	return http.ListenAndServe(addr, t)
}

func (t *Router) SetPrefix(prefix string) {
	t.prefix = prefix
	t.mainHandlerWrapper.SetPrefix(prefix)
}

func (t *Router) Use(handlerChain ...HandleFunc) {
	t.handlerChain = append(t.handlerChain, handlerChain...)
}
