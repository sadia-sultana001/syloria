package handler

import (
	"net/http"
	"syloria-demo/database"
	"syloria-demo/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	util.SendData(w, database.ProductList, 200)
}
