package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	Create(vehicle *Vehicle) (err error)

	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	GetAverageSpeedByBrand(brand string) (speed float64, err error)

	GetAverageCapacityByBrand(brand string) (capacity int, err error)

	UpdateSpeedById(id int, newSpeed float64) (v Vehicle, err error)

	UpdateFuelTypeById(id int, fuelType string) (v Vehicle, err error)

	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)

	FindByTransmission(transmission string) (v map[int]Vehicle, err error)

	Delete(id int) (err error)

	FindByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) (v map[int]Vehicle, err error)

	FindByWeightRange(minWeight float64, maxWeight float64) (v map[int]Vehicle, err error)
}
