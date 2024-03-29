// Code generated by yc@{{.CommandLineVersion}}. DO NOT EDIT.
package {{.PackageName}}

import (
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/cos"
	"github.com/juxuny/yc/env"
)

type _clientEnv struct {
	Entrypoint []string
	AccessKey  string
	Secret     string
}

type _cosClientEnv struct {
	Host    	  string
	RouteConfigId string
	AccessKey     string
	Secret        string
}

var clientEnv = _clientEnv{}
var cosClientEnv = _cosClientEnv{}

func init() {
	env.Init(&clientEnv, true, "{{.ServiceName|upper}}")
	env.Init(&cosClientEnv, true, "COS")
	initClient := &client{
		Service: Name,
	}
	if cosClientEnv.RouteConfigId == "" {
		initClient.EntrypointDispatcher = yc.NewRandomEntrypointDispatcher(clientEnv.Entrypoint)
	} else {
		initClient.EntrypointDispatcher = cos.NewEntrypointDispatcher(Name, cos.Options{
			CosEntrypoint: cosClientEnv.Host,
			ConfigId:      cosClientEnv.RouteConfigId,
			AccessKey:     cosClientEnv.AccessKey,
			Secret:        cosClientEnv.Secret,
		})
	}
	if clientEnv.AccessKey != "" && clientEnv.Secret != "" {
		initClient.signHandler = yc.NewDefaultSignHandler(clientEnv.AccessKey, clientEnv.Secret)
	}
	DefaultClient = initClient
}
