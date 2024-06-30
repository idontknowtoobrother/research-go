package main

import (
	"net"

	handler "github.com/idontknowtoobrother/kitchen/services/orders/handler/orders"
	"github.com/idontknowtoobrother/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("failed to listen", "error", err)
	}

	grpcServer := grpc.NewServer()

	// register gRPC services here
	ordersService := service.NewOrderService()
	handler.NewGRPCOrderService(grpcServer, ordersService)

	log.Print("gRPC server listening", "address", s.addr)

	return grpcServer.Serve(listener)
}
