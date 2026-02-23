package cmd

import (
	"fmt"
	"os"
	"syloria-demo/config"
	"syloria-demo/infra/db"
	"syloria-demo/product"
	"syloria-demo/repo"
	"syloria-demo/rest"

	producthandler "syloria-demo/rest/handler/product"

	usrHandler "syloria-demo/rest/handler/user"
	middleware "syloria-demo/rest/middlewares"
	"syloria-demo/user"
)

func Serve() {

	cnf := config.GetConfig()

	// fmt.Printf("&+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domain
	usrsvc := user.NewService(userRepo)
	prdSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := producthandler.NewHandler(middlewares, prdSvc)
	userHandler := usrHandler.NewHandler(cnf, usrsvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)

	server.Start()

}
