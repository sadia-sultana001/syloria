package product

import (
	"fmt"
	"net/http"
	"strconv"
	"syloria-demo/util"
)

func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	err = h.productRepo.Delete(pid)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	util.SendDate(w, http.StatusOK, "Successfully deleted product")

}
