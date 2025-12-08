package middleware

import (
	"log"
	"net/http"
)

func AnotherMidWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Another middleware")
		next.ServeHTTP(w, r)
	})
}
