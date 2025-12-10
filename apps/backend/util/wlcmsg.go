package util

import (
	"fmt"
	"net/http"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!\n")
}
