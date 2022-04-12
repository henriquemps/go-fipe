package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serverUp() {

	log.Println("Server started 🚀")

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
