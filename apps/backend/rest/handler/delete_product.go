package handler

import (
	"net/http"
	"strconv"
	"syloria-demo/database"
	"syloria-demo/util"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	database.Detele(pid)
	util.SendDate(w, "Successfully deleted product", 201)

}
