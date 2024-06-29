package main

import (
	"context"

	pb "github.com/idontknowtoobrother/omsv2-common/api"
)

type OrderService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context.Context) error
}
