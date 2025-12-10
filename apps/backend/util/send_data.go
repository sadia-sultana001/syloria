package util

import (
	"encoding/json"
	"net/http"
)

func SendDate(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(msg)
}
