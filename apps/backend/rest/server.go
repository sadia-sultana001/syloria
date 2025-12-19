package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"syloria-demo/config"
	"syloria-demo/util"

	"syloria-demo/rest/handler/product"
	"syloria-demo/rest/handler/review"
	"syloria-demo/rest/handler/user"
	middleware "syloria-demo/rest/middlewares"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	cnf config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(util.WelcomeMessage))

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	server.productHandler.ResisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on: ", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error staring the server", err)
		os.Exit(1)
	}
}
