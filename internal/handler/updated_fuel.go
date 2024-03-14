package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"slices"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *VehicleDefault) UpdateFuelTypeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

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

		var body VehicleJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if !slices.Contains([]string{"biodiesel", "diesel", "gas", "gasoline"}, body.FuelType) {
			response.Error(w, http.StatusBadRequest, "Invalid Fuel Type")
			return
		}

		vehicle, err := h.service.UpdateFuelTypeById(id, body.FuelType)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
			}

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicle updated",
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
