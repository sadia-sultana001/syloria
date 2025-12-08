package cmd

import (
	"fmt"
	"net/http"
	"syloria-demo/middleware"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!\n")
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", WelcomeMessage)

	manager := middleware.NewManager()

	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error staring the server", err)
	}
}
