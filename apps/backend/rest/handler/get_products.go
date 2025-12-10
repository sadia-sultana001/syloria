package handler

import (
	"net/http"
	"syloria-demo/database"
	"syloria-demo/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SendDate(w, database.List(), 200)
}
