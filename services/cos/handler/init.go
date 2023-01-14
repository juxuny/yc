package handler

import "github.com/juxuny/yc/router"

func init() {
	router.AddIgnoreAuthPath("/api/cos/login")
	router.AddIgnoreAuthPath("/api/cos/health")
	router.AddIgnoreAuthPath("/api/cos/list-all-value-by-config-id")
	router.AddOpenApiPath("/api/cos/list-all-value-by-config-id")
}
