package product

import (
	"net/http"
	"strconv"
	"syloria-demo/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	product, err := h.svc.Get(pid)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendDate(w, http.StatusOK, product)

}
