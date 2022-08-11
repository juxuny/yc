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

func AddIgnoreAuthPath(pathList ...string) {
	defaultRouter.AddIgnoreAuthPath(pathList...)
}

func AddOpenApiPath(pathList ...string) {
	defaultRouter.AddOpenApiPath(pathList...)
}

func CheckIfIgnoreAuth(path string) bool {
	return defaultRouter.CheckIfIgnoreAuth(path)
}

func CheckIfOpenApi(path string) bool {
	return defaultRouter.CheckIfOpenApi(path)
}
