package cmd

import (
	"fmt"
	"net/http"
	"syloria-demo/global_router"
	"syloria-demo/handler"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!\n")
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", WelcomeMessage)

	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))

	mux.Handle("POST /products/{productID}", http.HandlerFunc(handler.GetProductByID))

	fmt.Println("Server running on: 8080")

	globalRout := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRout)
	if err != nil {
		fmt.Println("Error staring the server", err)
	}
}
