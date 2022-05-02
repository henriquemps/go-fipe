package main

/**
 * Models
 *
 * Modelos para suporte de query vindas de metodo POST
 */

type QueryBrand struct {
	ReferenceId   int
	TypeVehicleId int
}

type QueryModel struct {
	TypeVehicleId int
	ReferenceId   int
	BrandId       int
}

type QueryYear struct {
	TypeVehicleId int
	ReferenceId   int
	BrandId       int
	ModelId       int
}

type QueryInfoVehicle struct {
	TypeVehicleId int
	ReferenceId   int
	BrandId       int
	ModelId       int
	YearModel     int
	CodeTypeFuel  int
	TypeVehicle   string
}
