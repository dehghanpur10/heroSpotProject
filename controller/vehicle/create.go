package vehicle

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/vehicleService"
)

// Create
// @Summary create a new vehicle for user
// @Description this endpoint creates a new vehicle for user
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param vehicle body models.Vehicle true "vehicle info"
// @Success 201 {object} models.Vehicle
// @Failure 400 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/vehicles [Post]
func Create(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	var vehicle models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		fmt.Printf("createVehicelHandler - decode - %v", err)
		lib.HttpError500(w)
		return
	}

	validate := validator.New()
	err = validate.Struct(vehicle)
	if err != nil {
		fmt.Printf("createVehicelHandler - validate - %v", err)
		lib.HttpError400(w, "all fields should be send")
		return
	}

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("createVehicelHandler - dynamoDb connect - %v", err)
		lib.HttpError500(w)
		return
	}

	serviceVehicle := vehicleService.New(db)
	err = serviceVehicle.Create(vehicle)
	if err != nil {
		fmt.Printf("createVehicelHandler - service - %v", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(vehicle)
	if err != nil {
		fmt.Printf("createVehicelHandler - marshal - %v", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusCreated, result)
}
