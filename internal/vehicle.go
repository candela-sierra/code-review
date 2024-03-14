package internal

func NewVehicle(id int, brand string, model string, registration string, color string, fabricationYear int,
	capacity int, maxSpeed float64, fuelType string, transmission string, weight float64, height float64,
	length float64, width float64) *Vehicle {
	return &Vehicle{id, VehicleAttributes{brand, model, registration, color, fabricationYear, capacity, maxSpeed, fuelType, transmission,
		weight, Dimensions{height, length, width}}}
}

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

func (vh Vehicle) EqualYear(year int) bool {
	return vh.FabricationYear == year
}

func (vh Vehicle) EqualColor(color string) bool {
	return vh.Color == color
}

func (vh Vehicle) EqualBrand(brand string) bool {
	return vh.Brand == brand
}

func (vh Vehicle) EqualFuelType(fuelType string) bool {
	return vh.FuelType == fuelType
}

func (vh Vehicle) EqualTransmission(transmission string) bool {
	return vh.Transmission == transmission
}

func (vh Vehicle) YearBetween(startYear int, endYear int) bool {
	return vh.FabricationYear > startYear && vh.FabricationYear < endYear
}

func (vh Vehicle) WidthBetween(minWidth float64, maxWidth float64) bool {
	return vh.Width > minWidth && vh.Width < maxWidth
}

func (vh Vehicle) HeightBetween(minHeight float64, maxHeight float64) bool {
	return vh.Height > minHeight && vh.Height < maxHeight
}

func (vh Vehicle) WeightBetween(minweight float64, maxWeight float64) bool {
	return vh.Weight > minweight && vh.Weight < maxWeight
}
