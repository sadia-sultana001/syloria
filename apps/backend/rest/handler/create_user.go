package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syloria-demo/database"
	"syloria-demo/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()
	util.SendDate(w, createdUser, http.StatusCreated)
}
