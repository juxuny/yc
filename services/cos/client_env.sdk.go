package cos

import "github.com/juxuny/yc/env"

type _clientEnv struct {
	Entrypoint string
}

var clientEnv = _clientEnv{}

func init() {
	env.Init(&clientEnv, true, "COS")
	DefaultClient = NewClient(clientEnv.Entrypoint)
}
