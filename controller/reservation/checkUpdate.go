package reservation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/service/reservationService"
)

// CheckUpdateReservationController
// @Summary checking  possibility  for update time
// @Description this endpoint will check  possibility for update time
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param reservation_id path string true "reservation ID"
// @Success 200 {object} models.Reservation
// @Failure 404 {object} lib.ErrorResponse
// @Failure 422 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/reservations/{reservation_id}/update [Get]
func CheckUpdateReservationController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("CheckUpdateReservationController - %v\n", err)
		lib.HttpError500(w)
		return
	}

	service := reservationService.New(db)
	vars := mux.Vars(r)
	reservation, err := service.CheckUpdateReservation(vars["reservation_id"])
	if err != nil {
		fmt.Printf("CheckUpdateReservationController - %v\n", err)
		if errors.Is(err, lib.ErrNotFound) {
			lib.HttpError404(w, "this reservation not found.")
			return
		}
		lib.HttpError500(w)
		return
	}

	if reservation.UpdatePossible {
		result, err := json.Marshal(reservation)
		if err != nil {
			fmt.Printf("CheckUpdateReservationController - marshal - %v\n", err)
			lib.HttpError500(w)
			return
		}
		lib.HttpSuccessResponse(w, http.StatusOK, result)
	} else {
		fmt.Printf("CheckUpdateReservationController - reservation can not be update\n")
		lib.HttpError422(w, "this reservation can not be update")
	}
}
