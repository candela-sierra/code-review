package repository

import (
	"app/internal"
	"errors"
)

func (repository *VehicleMap) Create(vehicle *internal.Vehicle) error {
	_, err := repository.FindById(vehicle.Id)

	if !errors.Is(internal.ErrVehicleNotFound, err) {
		return internal.ErrAlreadyExists
	}
	repository.db[vehicle.Id] = (*vehicle)
	return nil
}
