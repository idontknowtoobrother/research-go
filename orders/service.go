package main

import (
	"context"

	common "github.com/idontknowtoobrother/omsv2-common"
	pb "github.com/idontknowtoobrother/omsv2-common/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}

	mergeItemQuantities(&p.Items)
	logger.Info("validate order", "merged-items", p.Items)
	// validate with the stock service

	return nil
}

func mergeItemQuantities(items *[]*pb.ItemsWithQuantity) {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range *items {
		found := false
		for _, finalItem := range merged {
			if item.ID == finalItem.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	items = &merged
}
