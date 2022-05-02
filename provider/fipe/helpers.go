package fipe

const (
	VehicleCar        int = 1
	VehicleMotorcycle int = 2
	VehicleTruck      int = 3
)

func getTypeVehicle(typeId int) string {

	switch typeId {
	case VehicleCar:
		return "carro"
	case VehicleMotorcycle:
		return "moto"
	case VehicleTruck:
		return "caminhao"
	}

	return "carro"
}
