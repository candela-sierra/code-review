package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
)

func (h *VehicleDefault) FindByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		heights := strings.Split(r.URL.Query().Get("height"), "-")
		if len(heights) != 2 {
			response.Error(w, http.StatusBadRequest, "Invalid Height")
			return
		}
		minHeight, err := strconv.ParseFloat(heights[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Min Height")
			return
		}
		maxHeight, err := strconv.ParseFloat(heights[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Max Height")
			return
		}

		widths := strings.Split(r.URL.Query().Get("width"), "-")
		if len(heights) != 2 {
			response.Error(w, http.StatusBadRequest, "Invalid Width")
			return
		}
		minWidth, err := strconv.ParseFloat(widths[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Min Width")
			return
		}
		maxWidth, err := strconv.ParseFloat(widths[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Max Width")
			return
		}

		v, err := h.service.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
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
