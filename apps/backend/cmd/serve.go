package cmd

import (
	"syloria-demo/config"
	"syloria-demo/rest"
	"syloria-demo/rest/handler/product"
	"syloria-demo/rest/handler/review"
	"syloria-demo/rest/handler/user"
	middleware "syloria-demo/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)

	server.Start()

}
