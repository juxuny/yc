package router

var defaultRouter = NewRouter(
	"api",
	LoggerHandler,
	RecoverHandler,
	TraceHandler,
)

func Use(handlerChain ...HandleFunc) {
	defaultRouter.Use(handlerChain...)
}

func Register(groupName string, handler interface{}) error {
	return defaultRouter.Register(groupName, handler)
}

func SetPrefix(prefix string) {
	defaultRouter.SetPrefix(prefix)
}

func Start(addr string) error {
	return defaultRouter.Start(addr)
}
