package product

import (
	"fmt"
	"net/http"
	"syloria-demo/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	productList, err := h.svc.List()
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	util.SendDate(w, http.StatusOK, productList)
}
