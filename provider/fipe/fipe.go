package fipe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const VEHICLE_CAR int = 1
const VEHICLE_MOTORCYCLE int = 2
const VEHICLE_TRUCK int = 3

func GetReferences() []Reference {

	var references []Reference

	resource := fmt.Sprintf(os.Getenv("URL_BASE_FIPE"), "ConsultarTabelaDeReferencia")

	resp, err := http.Post(resource, "application/json", nil)

	if err != nil {
		resp.Body.Close()
		log.Fatalln("Error: %s", err)

		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&references); err != nil {
		resp.Body.Close()
		log.Fatalln("Error: %s", err)

		return nil
	}

	resp.Body.Close()

	return references
}

func GetBrands(referenceId int, typeVehicleId int) []Brand {

	var brands []Brand

	resource := fmt.Sprintf(os.Getenv("URL_BASE_FIPE"), "ConsultarMarcas")
	payload, _ := json.Marshal(map[string]int{
		"codigoTabelaReferencia": referenceId,
		"codigoTipoVeiculo":      typeVehicleId,
	})

	resp, err := http.Post(resource, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		resp.Body.Close()

		return []Brand{}
	}

	if err := json.NewDecoder(resp.Body).Decode(&brands); err != nil {
		resp.Body.Close()

		return []Brand{}
	}

	resp.Body.Close()

	return brands
}

func GetModels(referenceId int, typeVehicleId int, brandId int) []map[string]interface{} {

	var model Model

	resource := fmt.Sprintf(os.Getenv("URL_BASE_FIPE"), "ConsultarModelos")
	payload, _ := json.Marshal(map[string]int{
		"codigoTabelaReferencia": referenceId,
		"codigoTipoVeiculo":      typeVehicleId,
		"codigoMarca":            brandId,
	})

	resp, err := http.Post(resource, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		resp.Body.Close()

		return make([]map[string]interface{}, 0)
	}

	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		resp.Body.Close()

		return make([]map[string]interface{}, 0)
	}

	resp.Body.Close()

	return model.Modelos
}

func GetYears(referenceId int, typeVehicleId int, brandId int, modelId int) []Year {

	var years []Year

	resource := fmt.Sprintf(os.Getenv("URL_BASE_FIPE"), "ConsultarAnoModelo")
	payload, _ := json.Marshal(map[string]int{
		"codigoTabelaReferencia": referenceId,
		"codigoTipoVeiculo":      typeVehicleId,
		"codigoMarca":            brandId,
		"codigoModelo":           modelId,
	})

	resp, err := http.Post(resource, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		resp.Body.Close()

		return []Year{}
	}

	if err := json.NewDecoder(resp.Body).Decode(&years); err != nil {
		resp.Body.Close()

		return []Year{}
	}

	resp.Body.Close()

	return years
}
