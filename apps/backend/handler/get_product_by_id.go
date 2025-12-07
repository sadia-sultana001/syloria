package handler

import (
	"net/http"
	"strconv"
	"syloria-demo/database"
	"syloria-demo/util"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Product doesn't match.", 404)

}
