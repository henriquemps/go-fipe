package main

import (
	"fmt"
	"net/http"
)

func registerRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GO Fipe ")
	})

	http.HandleFunc("/ano-mes", GetReferences)

	http.HandleFunc("/marcas", GetBrands)

	http.HandleFunc("/modelos", GetModels)

	http.HandleFunc("/anos-modelo", GetYearsByModel)
}
