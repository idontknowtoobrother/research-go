package service

import (
	"context"

	"github.com/idontknowtoobrother/kitchen/services/common/genproto/orders"
)

var ordersDatabase = make([]*orders.Order, 0) // in-memory store (for demo purposes)

type OrderService struct {
	// store
	store *[]*orders.Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		store: &ordersDatabase,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDatabase = append(ordersDatabase, order)
	return nil
}
