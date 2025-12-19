package product

import (
	"net/http"
	middleware "syloria-demo/rest/middlewares"
)

func (h *Handler) ResisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProducts),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProducts),
		),
	)
}
