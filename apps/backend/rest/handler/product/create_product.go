package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"syloria-demo/repo"
	"syloria-demo/util"
)

type reqCreatePro struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req reqCreatePro
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	createdPro, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendDate(w, http.StatusCreated, createdPro)

}
