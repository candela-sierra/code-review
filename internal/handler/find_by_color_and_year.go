package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *VehicleDefault) FindByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid year")
			return
		}

		color := chi.URLParam(r, "color")
		if color == "" {
			response.Error(w, http.StatusBadRequest, "Invalid color")
			return
		}

		v, err := h.service.FindByColorAndYear(color, year)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "Vehicles not found")
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
			}

			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
