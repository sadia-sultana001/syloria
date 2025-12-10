package rest

import (
	"net/http"
	"syloria-demo/rest/handler"
	middleware "syloria-demo/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /manage",
		manager.With(
			http.HandlerFunc(handler.Test),
			middleware.AnotherMidWare,
		),
	)

	mux.Handle("GET /route",
		manager.With(
			http.HandlerFunc(handler.Test),
		),
	)

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handler.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handler.CreateProduct),
		),
	)

	mux.Handle("GET /products/{productID}",
		manager.With(
			http.HandlerFunc(handler.GetProductByID),
		),
	)
}
