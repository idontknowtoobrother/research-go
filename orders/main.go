package main

import (
	"context"
	"net"
	"os"
	"time"

	log "github.com/charmbracelet/log"
	common "github.com/idontknowtoobrother/omsv2-common"

	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "OMSV2 Orders 🧊",
	})
)

func main() {

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("failed to connect to gRPC server", "error", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	logger.Info("started gRPC server", "address", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		logger.Fatal("failed to start gRPC server", "error", err)
	}

}
