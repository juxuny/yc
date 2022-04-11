package http

import (
	"context"
	"fmt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/router"
	logServer "github.com/juxuny/yc/services/log-server"
	"github.com/juxuny/yc/services/log-server/config"
	"github.com/juxuny/yc/services/log-server/handler"
	"net/http"
)

func Start(ctx context.Context) {
	r := router.NewRouter("/api")
	if err := r.Register(logServer.Name, &handler.Handler{}); err != nil {
		panic(err)
	}
	finished := make(chan bool, 1)
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Env.HttpPort), r); err != nil {
			panic(err)
		}
		finished <- true
	}()
	select {
	case <-ctx.Done():
		log.Info("canceled")
		return
	case <-finished:
		log.Info("http server finished")
	}
}
