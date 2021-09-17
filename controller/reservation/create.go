package reservation

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/reservationService"
)

// CreateReservationController
// @Summary create a new reservation for vehicle
// @Description this endpoint creates a new reservation for vehicle
// @Tags reservation
// @Accept  json
// @Produce  json
// @Param reservation body models.InputReservation true "vehicle info"
// @Success 201 {object} models.Reservation
// @Failure 400 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/reservations [Post]
func CreateReservationController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	var inputReservation models.InputReservation
	err := json.NewDecoder(r.Body).Decode(&inputReservation)
	if err != nil {
		fmt.Printf("CreateReservationController - decode - %v", err)
		lib.HttpError500(w)
		return
	}

	validate := validator.New()
	err = validate.Struct(inputReservation)
	if err != nil {
		fmt.Printf("CreateReservationController - validate - %v", err)
		lib.HttpError400(w, "all fields should be send")
		return
	}

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("CreateReservationController - %v", err)
		lib.HttpError500(w)
		return
	}

	service := reservationService.New(db)

	reservation, err := service.FetchReservationInfo(inputReservation)
	if err != nil {
		fmt.Printf("CreateReservationController - %v", err)
		lib.HttpErrorWith(w, err)
		return
	}

	err = service.Create(reservation)
	if err != nil {
		fmt.Printf("CreateReservationController - %v", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(reservation)
	if err != nil {
		fmt.Printf("CreateReservationController - marshal - %v", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusCreated, result)
}
