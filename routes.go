package main

import (
	"net/http"
)

func registerRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ResponseJSON(w, map[string]string{"API": "GO Fipe"})
	})

	http.HandleFunc("/ano-mes", GetReferences)

	http.HandleFunc("/marcas", GetBrands)

	http.HandleFunc("/modelos", GetModels)

	http.HandleFunc("/anos-modelo", GetYearsByModel)

	http.HandleFunc("/dados-gerais", GetInfosGenerals)
}
