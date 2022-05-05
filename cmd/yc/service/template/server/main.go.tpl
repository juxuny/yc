package main

import (
	"context"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/services/{{.ServiceDir}}/server/http"
	"github.com/juxuny/yc/services/{{.ServiceDir}}/server/rpc"
	"github.com/juxuny/yc/trace"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, serverCanceler := context.WithCancel(context.Background())
	go rpc.Start(ctx)
	go http.Start(ctx)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	{{.Lt}}-quit
	log.Info("received shutdown server signal ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	finished := make(chan bool)
	go func() {
		trace.InitContext()
		trace.Wait()
		finished {{.Lt}}- true
	}()
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case {{.Lt}}-finished:
		serverCanceler()
		log.Info("server shutdown gracefully")
	case {{.Lt}}-ctx.Done():
		log.Info("timeout of 15 seconds.")
	}
}
