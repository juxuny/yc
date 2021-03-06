package main

import (
	"context"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/services/log-server/handler"
	"github.com/juxuny/yc/services/log-server/server/http"
	"github.com/juxuny/yc/services/log-server/server/rpc"
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
	<-quit
	log.Info("received shutdown server signal ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	finished := make(chan bool)
	go func() {
		trace.InitContext()
		if err := handler.Flush(); err != nil {
			log.Error(err)
		}
		trace.Wait()
		finished <- true
	}()
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-finished:
		serverCanceler()
		log.Info("server shutdown gracefully")
	case <-ctx.Done():
		log.Info("timeout of 15 seconds.")
	}
}
