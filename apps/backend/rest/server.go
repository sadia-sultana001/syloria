package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"syloria-demo/config"
	"syloria-demo/util"

	middleware "syloria-demo/rest/middlewares"
)

func Start(cnf config.Config) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(util.WelcomeMessage))

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server running on: ", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error staring the server", err)
		os.Exit(1)
	}
}
