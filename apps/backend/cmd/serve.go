package cmd

import (
	"syloria-demo/config"
	"syloria-demo/repo"
	"syloria-demo/rest"
	"syloria-demo/rest/handler/product"
	"syloria-demo/rest/handler/user"
	middleware "syloria-demo/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)

	server.Start()

}
