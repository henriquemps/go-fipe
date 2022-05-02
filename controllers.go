package main

import (
	"encoding/json"
	"fmt"
	"go-fipe/provider/fipe"
	"net/http"
)

func GetReferences(w http.ResponseWriter, r *http.Request) {

	if ValidMethod(w, r, http.MethodGet) == false {
		ResponseJSON(w, []string{})

		return
	}

	references := fipe.GetReferences()

	ResponseJSON(w, references)
}

func GetBrands(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if ValidMethod(w, r, http.MethodPost) == false {
		ResponseJSON(w, []string{})

		return
	}

	var params QueryBrand

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Error decoder params", http.StatusUnprocessableEntity)
	}

	brands := fipe.GetBrands(params.ReferenceId, params.TypeVehicleId)

	ResponseJSON(w, brands)
}

func GetModels(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if ValidMethod(w, r, http.MethodPost) == false {
		ResponseJSON(w, []string{})

		return
	}

	var params QueryModel

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Error decoder params", http.StatusUnprocessableEntity)
	}

	models := fipe.GetModels(params.ReferenceId, params.TypeVehicleId, params.BrandId)

	ResponseJSON(w, models)
}

func GetYearsByModel(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if ValidMethod(w, r, http.MethodPost) == false {
		ResponseJSON(w, []string{})

		return
	}

	var params QueryYear

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Error decoder params", http.StatusUnprocessableEntity)
	}

	years := fipe.GetYears(params.ReferenceId, params.TypeVehicleId, params.BrandId, params.ModelId)

	ResponseJSON(w, years)
}

func GetInfosGenerals(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if ValidMethod(w, r, http.MethodPost) == false {
		ResponseJSON(w, []string{})

		return
	}

	var params QueryInfoVehicle

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, fmt.Sprintf("Error decoder params: %s", err), http.StatusInternalServerError)
	}

	info := fipe.GetInfosGenerals(params.ReferenceId, params.TypeVehicleId, params.BrandId, params.ModelId, params.YearModel, params.CodeTypeFuel)

	ResponseJSON(w, info)
}
