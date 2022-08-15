package config

import "github.com/juxuny/yc/env"

var Env = struct {
	HttpPort     int    `env:"20080"`
    Prefix       string `env:"/api"`
}{}

func init() {
	env.Init(&Env, true)
}
