package reservation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/reservationService"
)

// UpdateReservationController
// @Summary update reservation time
// @Description this endpoint will update reservation (url reservation_id should be equal to body reservation_id)
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param reservation_id path string true "reservation ID"
// @Param reservation body models.InputReservation true "reservation info"
// @Success 200 {object} models.Reservation
// @Failure 404 {object} lib.ErrorResponse
// @Failure 400 {object} lib.ErrorResponse
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
	reservation, err := service.CheckUpdateReservation(vars["reservation_id"])
	if err != nil {
		fmt.Printf("UpdateReservationController - %v", err)
		if errors.Is(err, lib.ErrNotFound) {
			lib.HttpError404(w, "reservation not found")
			return
		}
		lib.HttpError500(w)
		return
	}

	if !reservation.UpdatePossible {
		fmt.Printf("UpdateReservationController - reservation can not be update")
		lib.HttpError422(w, "this reservation can not be update")
		return
	}

	var inputReservation models.InputReservation
	err = json.NewDecoder(r.Body).Decode(&inputReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - decode - %v", err)
		lib.HttpError400(w, "please enter correct body request")
		return
	}

	if inputReservation.Id != vars["reservation_id"] {
		fmt.Println("UpdateReservationController - No match url id and body id")
		lib.HttpError400(w, "url reservation_id must be equal to body reservation_id")
		return
	}

	validate := validator.New()
	err = validate.Struct(inputReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - validate - %v", err)
		lib.HttpError400(w, "all fields should be send")
		return
	}

	newReservation, err := service.FetchReservationInfo(inputReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - %v\n", err)
		if errors.Is(err, lib.ErrNotFound) {
			lib.HttpError404(w, "facility_id or vehicle_id not found")
			return
		}
		lib.HttpError500(w)
		return
	}

	err = service.CreateReservation(newReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - %v\n", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(newReservation)
	if err != nil {
		fmt.Printf("UpdateReservationController - marshal - %v\n", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, result)

}
