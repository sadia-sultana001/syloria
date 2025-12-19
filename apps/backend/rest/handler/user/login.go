package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syloria-demo/config"
	"syloria-demo/database"
	"syloria-demo/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JWTsecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendDate(w, accessToken, http.StatusCreated)

}
