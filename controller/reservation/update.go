package reservation

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/reservationService"
)

// UpdateReservationController
// @Summary update reservation time
// @Description this endpoint will update reservation time
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param reservation_id path string true "reservation ID"
// @Param reservation body models.InputReservation true "reservation info"
// @Success 200 {object} models.Reservation
// @Failure 404 {object} lib.ErrorResponse
// @Failure 422 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/reservations/{reservation_id}/update [Put]
func UpdateReservationController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("UpdateReservationController - %v", err)
		lib.HttpError500(w)
		return
	}

	service := reservationService.New(db)
	vars := mux.Vars(r)
	reservation, err := service.CheckUpdate(vars["reservation_id"])
	if err != nil {
		fmt.Printf("UpdateReservationController - %v", err)
		lib.HttpErrorWith(w, err)
		return
	}

	if !reservation.UpdatePossible {
		fmt.Printf("reservation can not be update")
		lib.HttpError422(w, "this reservation can not be update")
		return
	}

	var inputReservation models.InputReservation
	err = json.NewDecoder(r.Body).Decode(&inputReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - decode - %v", err)
		lib.HttpError500(w)
		return
	}

	newReservation, err := service.FetchReservationInfo(inputReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - %v", err)
		lib.HttpErrorWith(w, err)
		return
	}

	err = service.Create(newReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - %v", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(newReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - marshal - %v", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusCreated, result)

}
