package reservation

import (
	"fmt"
	"net/http"
	"spotHeroProject/lib"
)
// CheckUpdate
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
func CheckUpdate(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Println("check update  controller - connect to dynamoDb: ", err)
		lib.HttpError500(w)
		return
	}
	_ = db
	// Todo add check update reservation service
}
