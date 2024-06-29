package main

import (
	"errors"
	"net/http"

	common "github.com/idontknowtoobrother/omsv2-common"
	pb "github.com/idontknowtoobrother/omsv2-common/api"
)

type handler struct {
	// gateway
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	logger.Info("received request", "method", r.Method, "path", r.URL.Path)
	// time.Sleep(500 * time.Millisecond)
	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJson(r, &items); err != nil {
		logger.Error("failed to read request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := validateItems(&items)
	if err != nil {
		logger.Warn("failed to validate items", "error", err)
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	logger.Info("redirecting", "user-ip", r.RemoteAddr, "path", r.URL.Path)
	orderCreated, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	if err != nil {
		logger.Error("failed to create order", "error", err)
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, orderCreated)
	logger.Info("order created", "order", orderCreated)
}

func validateItems(items *[]*pb.ItemsWithQuantity) error {
	if len(*items) == 0 {
		return common.ErrNoItems
	}

	for _, item := range *items {
		if item.ID == "" {
			return errors.New("item ID is required")
		}

		if item.Quantity <= 0 {
			return errors.New("item must have a valid quantity")
		}
	}

	return nil
}
