package reservation

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/facilityService"
	"spotHeroProject/service/reservationService"
	"spotHeroProject/service/vehicleService"
	"testing"
)

func TestCheckUpdateReservationControllerSuccess(t *testing.T) {
	// Arrange
	vehicle := models.Vehicle{
		Id: "testVehicleId",
	}
	facility := models.Facility{
		Id: "testFacilityId",
	}
	reservationInput := models.InputReservation{
		Id:             "testReservationId",
		Facility:       "testFacilityId",
		ParkedVehicle:  "testVehicleId",
		UpdatePossible: true,
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	}

	db, err := lib.GetDynamoDB()
	require.NoError(t, err)
	serviceOfVehicle := vehicleService.New(db)
	err = serviceOfVehicle.CreateVehicle(vehicle)
	require.NoError(t, err)
	serviceOfFacility := facilityService.New(db)
	err = serviceOfFacility.CreateFacilityService(facility)
	require.NoError(t, err)
	serviceOfReservation := reservationService.New(db)
	reservation, err := serviceOfReservation.FetchReservationInfo(reservationInput)
	require.NoError(t, err)
	err = serviceOfReservation.CreateReservation(reservation)
	require.NoError(t, err)

	expectedStatus := 200
	expectedReservation := models.Reservation{
		Id:             "testReservationId",
		Facility:       facility,
		ParkedVehicle:  vehicle,
		UpdatePossible: true,
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", CheckUpdateReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation/testReservationId/update", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var output models.Reservation
	err = json.Unmarshal(rr.Body.Bytes(), &output)
	require.NoError(t, err)
	assert.Equal(t, expectedReservation, output)

	err = serviceOfFacility.DeleteFacilityService(facility.Id)
	require.NoError(t, err)
	err = serviceOfVehicle.DeleteVehicle(vehicle.Id)
	require.NoError(t, err)
	err = serviceOfReservation.DeleteReservationService(reservationInput.Id)
	require.NoError(t, err)
}

func TestCheckUpdateReservationControllerNotAllow(t *testing.T) {
	// Arrange
	vehicle := models.Vehicle{
		Id: "testVehicleId",
	}
	facility := models.Facility{
		Id: "testFacilityId",
	}
	reservationInput := models.InputReservation{
		Id:             "testReservationId",
		Facility:       "testFacilityId",
		ParkedVehicle:  "testVehicleId",
		UpdatePossible: false,
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	}

	db, err := lib.GetDynamoDB()
	require.NoError(t, err)
	serviceOfVehicle := vehicleService.New(db)
	err = serviceOfVehicle.CreateVehicle(vehicle)
	require.NoError(t, err)
	serviceOfFacility := facilityService.New(db)
	err = serviceOfFacility.CreateFacilityService(facility)
	require.NoError(t, err)
	serviceOfReservation := reservationService.New(db)
	reservation, err := serviceOfReservation.FetchReservationInfo(reservationInput)
	require.NoError(t, err)
	err = serviceOfReservation.CreateReservation(reservation)
	require.NoError(t, err)

	expectedStatus := 422
	expectedError := lib.ErrorResponse{Code: 422, Title: "unprocessable input", Description: "this reservation can not be update"}

	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", CheckUpdateReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation/testReservationId/update", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = serviceOfFacility.DeleteFacilityService(facility.Id)
	require.NoError(t, err)
	err = serviceOfVehicle.DeleteVehicle(vehicle.Id)
	require.NoError(t, err)
	err = serviceOfReservation.DeleteReservationService(reservationInput.Id)
	require.NoError(t, err)
}

func TestCheckUpdateReservationControllerFailOnGetDynamoDB(t *testing.T) {
	// Arrange
	defer func() {
		lib.AWS_REGION = "us-west-2"
	}()
	lib.AWS_REGION = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", CheckUpdateReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation/testReservationId/update", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestCheckUpdateReservationControllerFailOnCheckService(t *testing.T) {
	// Arrange
	defer func() {
		lib.RESERVATION_TABLE_NAME = "ReservationSpot"
	}()
	lib.RESERVATION_TABLE_NAME = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", CheckUpdateReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation/testReservationId/update", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestCheckUpdateReservationControllerFailOnNotFound(t *testing.T) {
	// Arrange
	expectedStatus := 404
	expectedError := lib.ErrorResponse{Code: 404, Title: "Not found", Description: "this reservation not found."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", CheckUpdateReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation/testReservationId/update", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	fmt.Println(rr.Body.String())
	err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}