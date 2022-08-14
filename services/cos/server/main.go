// Code generated by yc@v0.0.1
package main

import (
	"context"
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/services/cos/server/http"
	"github.com/juxuny/yc/trace"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serveCommand struct {
}

func (t *serveCommand) Prepare(cmd *cobra.Command) {
}

func (t *serveCommand) InitFlag(cmd *cobra.Command) {
}

func (t *serveCommand) BeforeRun(cmd *cobra.Command) {
}

func (t *serveCommand) Run() {
	ctx, serverCanceler := context.WithCancel(context.Background())
	defer serverCanceler()
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

var rootCommand = &cobra.Command{}

func main() {
	rootCommand.AddCommand(cmd.NewCommandBuilder("serve", &serveCommand{}).Build())
	if err := rootCommand.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
