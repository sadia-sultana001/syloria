package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syloria-demo/product"
	"syloria-demo/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	util.SendDate(w, newProduct, 201)

}
