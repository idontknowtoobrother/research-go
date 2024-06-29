package main

import (
	"context"

	pb "github.com/idontknowtoobrother/omsv2-common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrderService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	logger.Info("new order received!", "orders", p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
