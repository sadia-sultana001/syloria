package cmd

import (
	"syloria-demo/config"
	"syloria-demo/rest"
	"syloria-demo/rest/handler/product"
	"syloria-demo/rest/handler/review"
	"syloria-demo/rest/handler/user"
)

func Serve() {

	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)

	server.Start(cnf)

}
