package fipe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

	for i, item := range years {

		slice := strings.Split(item.Value, "-")

		ano := slice[0]
		combustivel := slice[1]

		years[i].Ano, _ = strconv.Atoi(ano)
		years[i].Combustivel, _ = strconv.Atoi(combustivel)
	}

	resp.Body.Close()

	return years
}

func GetInfosGenerals(referenceId int, typeVehicleId int, brandId int, modelId int, yearModel int, codeTypeFuel int) InfoVehicle {

	var info InfoVehicle

	resource := fmt.Sprintf(os.Getenv("URL_BASE_FIPE"), "ConsultarValorComTodosParametros")
	payload, _ := json.Marshal(map[string]interface{}{
		"codigoTabelaReferencia": referenceId,
		"codigoTipoVeiculo":      typeVehicleId,
		"codigoMarca":            brandId,
		"codigoModelo":           modelId,
		"anoModelo":              yearModel,
		"codigoTipoCombustivel":  codeTypeFuel,
		"tipoVeiculo":            getTypeVehicle(typeVehicleId),
		"tipoConsulta":           "tradicional",
	})

	resp, err := http.Post(resource, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		resp.Body.Close()

		return info
	}

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		resp.Body.Close()

		return info
	}

	resp.Body.Close()

	return info
}
