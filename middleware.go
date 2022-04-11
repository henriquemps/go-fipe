package main

import (
	"net/http"
)

func ValidMethod(w http.ResponseWriter, r *http.Request, method string) bool {

	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return false
	}

	return true
}
