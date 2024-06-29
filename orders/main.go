package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	defer grpcServer.Stop()

	// START Graceful shutdown (need to learn more about this)
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-term
		logger.Info("gracefully shutting down server...")
		if err := l.Close(); err != nil {
			logger.Fatal("failed to shutdown server", "error", err)
		}
	}()
	// END Graceful shutdown

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	logger.Info("started gRPC server", "address", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		logger.Fatal("failed to start gRPC server", "error", err)
	}

}
