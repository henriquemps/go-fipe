package main

/**
 * Models
 *
 * Modelos para suporte de query vindas de method POST
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
