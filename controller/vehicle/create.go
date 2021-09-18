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

// CreateVehicleController
// @Summary create a new vehicle for user
// @Description this endpoint creates a new vehicle for user
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param vehicle body models.Vehicle true "vehicle info"
// @Success 201 {object} models.Vehicle "vehicle created successfully"
// @Failure 400 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/vehicles [Post]
func CreateVehicleController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	var vehicle models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		fmt.Printf("CreateVehicleController - decode - %v", err)
		lib.HttpError400(w, "please enter correct body request")
		return
	}

	validate := validator.New()
	err = validate.Struct(vehicle)
	if err != nil {
		fmt.Printf("CreateVehicleController - validate - %v", err)
		lib.HttpError400(w, "all fields should be send")
		return
	}

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("CreateVehicleController - %v", err)
		lib.HttpError500(w)
		return
	}

	serviceVehicle := vehicleService.New(db)
	err = serviceVehicle.Create(vehicle)
	if err != nil {
		fmt.Printf("CreateVehicleController - %v", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(vehicle)
	if err != nil {
		fmt.Printf("CreateVehicleController - marshal - %v", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusCreated, result)
}
