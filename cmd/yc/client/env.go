package client

import "github.com/juxuny/yc/env"

var Env = struct {
	Gopath string
	YcHome string
}{}

func init() {
	env.Init(&Env, true)
}
