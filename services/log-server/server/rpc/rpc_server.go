package rpc

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/log"
	logServer "github.com/juxuny/yc/services/log-server"
	"github.com/juxuny/yc/services/log-server/handler"
	"google.golang.org/grpc"
	"net"
	"os"
)

var rpcServer *grpc.Server

func Start(ctx context.Context) {
	listenAddress := fmt.Sprintf(":%d", yc.DefaultRpcPort)
	ln, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	log.Info("start rpc ")
	opts := make([]grpc.ServerOption, 0)
	rpcServer = grpc.NewServer(opts...)
	logServer.RegisterLogServerServer(rpcServer, &handler.Handler{})
	finished := make(chan bool, 1)
	go func() {
		if err := rpcServer.Serve(ln); err != nil {
			log.Error("failed to serve:", err)
			rpcServer = nil
			finished <- true
		}
	}()
	select {
	case <-ctx.Done():
		Stop()
		return
	case <-finished:
		return
	}
}

func Stop() {
	if rpcServer == nil {
		return
	}
	rpcServer.Stop()
}
