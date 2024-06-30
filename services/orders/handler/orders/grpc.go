package handler

import (
	"context"

	"github.com/idontknowtoobrother/kitchen/services/common/genproto/orders"
	"github.com/idontknowtoobrother/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGRPCHandler struct {
	// service injection
	orderService types.OrderService
	// unimplemented UnimplementedOrderServiceServer
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrderService(grpcServer *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrderGRPCHandler{
		orderService: orderService,
	}

	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpcServer, gRPCHandler)
}

func (handler *OrderGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    1,
		CustomerID: 10,
		ProductID:  1,
		Quantity:   10,
	}

	err := handler.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Order created successfully",
	}

	return res, nil
}
