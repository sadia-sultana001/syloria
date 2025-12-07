package cmd

import (
	"fmt"
	"net/http"
	"syloria-demo/global_router"
	"syloria-demo/middleware"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!\n")
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", WelcomeMessage)

	manager := middleware.NewMessage()

	manager.Use(middleware.Hudai, middleware.Logger)

	initRoutes(mux, manager)

	//globalRouter := global_router.GlobalRouter(mux)

	globalRout := global_router.GlobalRouter(mux)

	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", globalRout)
	if err != nil {
		fmt.Println("Error staring the server", err)
	}
}
