package repository

import "app/internal"

func (repository *VehicleMap) UpdateSpeedById(id int, newSpeed float64) (v internal.Vehicle, err error) {
	currentVehicle, err := repository.FindById(id)
	if err != nil {
		return
	}

	currentVehicle.MaxSpeed = newSpeed
	repository.db[id] = currentVehicle
	v = currentVehicle
	return
}

func (repository *VehicleMap) UpdateFuelTypeById(id int, fuelType string) (v internal.Vehicle, err error) {
	currentVehicle, err := repository.FindById(id)
	if err != nil {
		return
	}

	currentVehicle.FuelType = fuelType
	repository.db[id] = currentVehicle
	v = currentVehicle
	return
}
