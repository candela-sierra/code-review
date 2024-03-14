package repository

import "app/internal"

func (r *VehicleMap) FindBy(condition func(internal.Vehicle) bool) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if condition(value) {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool { return v.EqualYear(year) && v.EqualColor(color) })
	return
}

func (r *VehicleMap) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool { return v.EqualBrand(brand) && v.YearBetween(startYear, endYear) })
	return
}

func (r *VehicleMap) FindByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool {
		return v.HeightBetween(minHeight, maxHeight) && v.WidthBetween(minWidth, maxWidth)
	})
	return
}

func (r *VehicleMap) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool { return v.EqualFuelType(fuelType) })
	return
}

func (r *VehicleMap) FindByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool { return v.EqualTransmission(transmission) })
	return
}

func (r *VehicleMap) FindByWeightRange(minweight float64, maxweight float64) (v map[int]internal.Vehicle, err error) {
	v, err = r.FindBy(func(v internal.Vehicle) bool {
		return v.WeightBetween(minweight, maxweight)
	})
	return
}

func (r *VehicleMap) FindById(id int) (internal.Vehicle, error) {
	v, ok := r.db[id]
	if !ok {
		return v, internal.ErrVehicleNotFound
	}
	return v, nil
}
