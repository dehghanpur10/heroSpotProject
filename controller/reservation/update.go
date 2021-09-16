package reservation

import (
	"fmt"
	"net/http"
	"spotHeroProject/lib"
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
	_ = db

	// Todo add check update reservation service

	// Todo add update reservation time
}
