package cmd

import (
	"fmt"
	"os"
	"syloria-demo/config"
	"syloria-demo/infra/db"
	"syloria-demo/repo"
	"syloria-demo/rest"
	"syloria-demo/rest/handler/product"
	"syloria-demo/rest/handler/user"
	middleware "syloria-demo/rest/middlewares"
)

func Serve() {

	cnf := config.GetConfig()

	dbCon, err := db.NweConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
