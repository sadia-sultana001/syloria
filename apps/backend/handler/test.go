package handler

import (
	"log"
	"net/http"
)

func Test(next http.ResponseWriter, r *http.Request) {
	log.Println("2nd print/middle e print- handler")
}
