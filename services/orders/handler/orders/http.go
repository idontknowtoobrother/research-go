package handler

import (
	"net/http"

	log "github.com/charmbracelet/log"
	"github.com/idontknowtoobrother/kitchen/services/common/genproto/orders"
	"github.com/idontknowtoobrother/kitchen/services/common/utils"
	"github.com/idontknowtoobrother/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}
	return handler
}

func (handler *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/orders", handler.CreateOrder)
}

func (handler *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := utils.DecodeRequest(r, &req)
	if err != nil {
		log.Info("Error decoding request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = handler.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		log.Info("Error creating order", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("Order created successfully", "order", order)
	utils.EncodeResponse(w, &orders.CreateOrderResponse{
		Status: "Order created successfully",
	})
}
