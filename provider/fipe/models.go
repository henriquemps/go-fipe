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
	Label string
	Value string
}
