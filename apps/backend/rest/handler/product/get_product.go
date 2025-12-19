package product

import (
	"net/http"
	"strconv"
	"syloria-demo/database"
	"syloria-demo/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product := database.Get(pid)
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendDate(w, product, 200)

}
