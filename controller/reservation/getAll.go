package reservation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/service/reservationService"
)

// GetAllReservationController
// @Summary Get the summary of all reservations
// @Description this endpoint Get the summary of all reservations
// @Tags reservation
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Reservation
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/reservations [Get]
func GetAllReservationController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Println("GetAllReservationController -  ", err)
		lib.HttpError500(w)
		return
	}
	service := reservationService.New(db)

	reservations, err := service.GetAllReservation()
	if err != nil {
		fmt.Println("GetAllReservationController -  ", err)
		lib.HttpError500(w)
	}

	result, err := json.Marshal(reservations)
	if err != nil {
		fmt.Printf("GetAllReservationController - Marshal - %v\n", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusOK, result)
}
