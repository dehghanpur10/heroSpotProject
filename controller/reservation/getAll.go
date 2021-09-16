package reservation

import (
	"fmt"
	"net/http"
	"spotHeroProject/lib"
)
// GetAll
// @Summary Get the summary of all reservations
// @Description this endpoint Get the summary of all reservations
// @Tags reservation
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Reservation
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/reservations [Get]
func GetAll(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Println("get all reservation controller - connect to dynamoDb: ", err)
		lib.HttpError500(w)
		return
	}
	_ = db
	// Todo add get all reservation service
}