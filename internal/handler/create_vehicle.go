package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func (handler *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		err = validateRequiredFields(bodyMap, "id", "brand", "model", "registration", "color", "fabrication_year", "capacity",
			"max_speed", "fuel_type", "transmission", "weight", "height", "length", "width")

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		var body VehicleJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		vehicle := internal.NewVehicle(body.ID, body.Brand, body.Model, body.Registration,
			body.Color, body.FabricationYear, body.Capacity, body.MaxSpeed, body.FuelType,
			body.Transmission, body.Weight, body.Height, body.Length, body.Width)

		err = handler.service.Create(vehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrAlreadyExists):
				response.Error(w, http.StatusConflict, "Vehicle already exists")
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
			}

			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehicle created",
			"data": VehicleJSON{
				ID:              vehicle.Id,
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Height:          vehicle.Height,
				Length:          vehicle.Length,
				Width:           vehicle.Width,
			},
		})
	}
}

func validateRequiredFields(bodyMap map[string]any, keys ...string) error {
	for _, key := range keys {
		if _, ok := bodyMap[key]; !ok {
			return fmt.Errorf("%s required", key)
		}
	}
	return nil
}
