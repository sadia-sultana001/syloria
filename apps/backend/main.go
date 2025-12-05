package main

import (
	"syloria-demo/cmd"
)

/* func corsMiddleWare(next http.Handler) http.Handler {

	enableCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	}
	handler := http.HandlerFunc(enableCors)
	return handler
} */

func main() {
	cmd.Serve()

}
