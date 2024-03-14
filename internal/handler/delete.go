package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Id")
			return
		}

		err = h.service.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
			}

			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
