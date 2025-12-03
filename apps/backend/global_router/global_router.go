package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleReq := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(handleReq)
}
