package main

import (
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	log "github.com/charmbracelet/log"
	common "github.com/idontknowtoobrother/omsv2-common"
	pb "github.com/idontknowtoobrother/omsv2-common/api"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "OMSV2 Gateway 🧊",
	})

	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		logger.Fatal("failed to connect to order service", "error", err)
	}
	defer conn.Close()

	logger.Info("connected to order service", "address", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)
	logger.Info("starting server", "port", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		logger.Fatal("failed to start server", "error", err)
	}
}
