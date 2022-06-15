package {{.PackageName}}

import (
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/env"
)

type _clientEnv struct {
	Entrypoint []string
}

var clientEnv = _clientEnv{}

func init() {
	env.Init(&clientEnv, true, "{{.ServiceName|upper}}")
	DefaultClient = NewClientWithDispatcher(yc.NewRandomEntrypointDispatcher(clientEnv.Entrypoint))
}
