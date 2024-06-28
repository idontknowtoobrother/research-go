package handlers

import (
	"fmt"
	"net/http"
)

	type Second struct{}

func (rcv Second) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "second response")
}
