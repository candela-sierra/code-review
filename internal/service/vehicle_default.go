package service

import (
	"app/internal"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{respository: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// respository is the repository that will be used by the service
	respository internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindAll()
	return
}

func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	return s.respository.Create(v)
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByColorAndYear(color, year)
	return
}

func (s *VehicleDefault) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByBrandAndYearRange(brand, startYear, endYear)
	return
}

func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (speed float64, err error) {
	speed, err = s.respository.GetAverageSpeedByBrand(brand)
	return
}

func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (capacity int, err error) {
	capacity, err = s.respository.GetAverageCapacityByBrand(brand)
	return
}

func (s *VehicleDefault) UpdateSpeedById(id int, newSpeed float64) (v internal.Vehicle, err error) {
	v, err = s.respository.UpdateSpeedById(id, newSpeed)
	return
}

func (s *VehicleDefault) UpdateFuelTypeById(id int, fuelType string) (v internal.Vehicle, err error) {
	v, err = s.respository.UpdateFuelTypeById(id, fuelType)
	return
}

func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByFuelType(fuelType)
	return
}

func (s *VehicleDefault) FindByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByTransmission(transmission)
	return
}

func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.respository.Delete(id)
	return
}

func (s *VehicleDefault) FindByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
	return
}

func (s *VehicleDefault) FindByWeightRange(minWeight float64, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.respository.FindByWeightRange(minWeight, maxWeight)
	return
}
