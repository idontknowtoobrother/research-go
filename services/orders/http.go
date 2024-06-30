package main

import (
	"net/http"

	handler "github.com/idontknowtoobrother/kitchen/services/orders/handler/orders"
	"github.com/idontknowtoobrother/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderhandler := handler.NewHttpOrdersHandler(orderService)

	orderhandler.RegisterRouter(router)

	log.Info("HTTP server listening", "address", s.addr)

	return http.ListenAndServe(s.addr, router)
}
