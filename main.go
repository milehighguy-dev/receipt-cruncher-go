package main

import (
	"fmt"
	"net/http"
	"github.com/milehighguy-dev/receipt-cruncher-go/pkg/controller"
)

func main() {
	controller.RegisterRoutes()
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}