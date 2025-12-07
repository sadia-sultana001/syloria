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

	manager := middleware.NewMessage()

	manager.Use(middleware.Hudai, middleware.Logger, middleware.CorsWithPreflight)

	initRoutes(mux, manager)

	//globalRout := middleware.CorsWithPreflight(mux)

	wrapMux := manager.With(mux)

	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", wrapMux)
	if err != nil {
		fmt.Println("Error staring the server", err)
	}
}
