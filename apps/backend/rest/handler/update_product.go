package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"syloria-demo/database"
	"syloria-demo/util"
)

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = pid
	database.Update(newProduct)
	util.SendDate(w, "Successfully updated product", 201)

}
