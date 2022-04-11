package main

import (
	"log"
	"net/http"
)

func serverUp() {

	log.Println("Server started 🚀")

	http.ListenAndServe(":8080", nil)
}
