package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"syloria-demo/repo"
	"syloria-demo/util"
)

type reqUpdatePro struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	var req reqUpdatePro
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          pid,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	util.SendDate(w, http.StatusOK, "Successfully updated product")

}
