package main

import (
	"net/http"

	"github.com/backend/middleware/handlers"
)

func main() {
	http.Handle("/first", handlers.First{})
	http.Handle("/second", handlers.Second{})

	http.ListenAndServe(":8080", nil)
}
