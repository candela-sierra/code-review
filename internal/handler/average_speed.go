package handler

import (
	"app/internal"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Invalid brand")
			return
		}

		speed, err := h.service.GetAverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "Vehicles not found")
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
			}

			return
		}

		data := make(map[string]any)
		data["average_speed"] = speed
		data["brand"] = brand
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
