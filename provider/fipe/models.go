package fipe

type Reference struct {
	Codigo int
	Mes    string
}

type Brand struct {
	Value string
	Label string
}

type Model struct {
	Anos    []map[string]string
	Modelos []map[string]interface{}
}

type Year struct {
	Label       string
	Value       string
	Ano         int
	Combustivel int
}

type InfoVehicle struct {
	AnoModelo        int
	Autenticacao     string
	CodigoFipe       string
	Combustivel      string
	DataConsulta     string
	Marca            string
	MesReferencia    string
	Modelo           string
	SiglaCombustivel string
	TipoVeiculo      int
	Valor            string
}
