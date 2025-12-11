package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syloria-demo/database"
	"syloria-demo/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credential.", http.StatusBadRequest)
		return
	}
	util.SendDate(w, usr, http.StatusCreated)

}
