package repository

import "app/internal"

func (repository *VehicleMap) GetAverageSpeedByBrand(brand string) (speed float64, err error) {
	speed = 0
	amount := 0.0
	for _, value := range repository.db {
		if value.EqualBrand(brand) {
			speed += value.MaxSpeed
			amount++
		}
	}

	if amount == 0 {
		err = internal.ErrVehicleNotFound
	}

	speed /= amount
	return
}

func (repository *VehicleMap) GetAverageCapacityByBrand(brand string) (capacity int, err error) {
	capacity = 0
	amount := 0
	for _, value := range repository.db {
		if value.EqualBrand(brand) {
			capacity += value.Capacity
			amount++
		}
	}

	if amount == 0 {
		err = internal.ErrVehicleNotFound
	}

	capacity /= amount
	return
}
