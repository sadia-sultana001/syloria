package handler

import (
	"net/http"
	"syloria-demo/product"
	"syloria-demo/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	util.SendDate(w, product.ProductList, 200)
}
