package reservation

import (
	"bytes"
	"encoding/json"
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

func createReservation(possibleUpdate bool) (models.Reservation, error) {
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
		UpdatePossible: possibleUpdate,
		Quote: models.Quote{
			Ends:   "8",
			Starts: "5",
		},
	}

	db, err := lib.GetDynamoDB()
	if err != nil {
		return models.Reservation{}, nil
	}
	serviceOfVehicle := vehicleService.New(db)
	err = serviceOfVehicle.CreateVehicle(vehicle)
	if err != nil {
		return models.Reservation{}, nil
	}
	serviceOfFacility := facilityService.New(db)
	err = serviceOfFacility.CreateFacilityService(facility)
	if err != nil {
		return models.Reservation{}, nil
	}
	serviceOfReservation := reservationService.New(db)
	reservation, err := serviceOfReservation.FetchReservationInfo(reservationInput)
	if err != nil {
		return models.Reservation{}, nil
	}
	err = serviceOfReservation.CreateReservation(reservation)
	if err != nil {
		return models.Reservation{}, nil
	}
	return reservation, nil
}

func deleteReservation() error {
	db, err := lib.GetDynamoDB()
	if err != nil {
		return err
	}
	serviceOfVehicle := vehicleService.New(db)
	serviceOfFacility := facilityService.New(db)
	serviceOfReservation := reservationService.New(db)

	err = serviceOfFacility.DeleteFacilityService("testFacilityId")
	if err != nil {
		return err
	}
	err = serviceOfVehicle.DeleteVehicle("testVehicleId")
	if err != nil {
		return err
	}
	err = serviceOfReservation.DeleteReservationService("testReservationId")
	if err != nil {
		return err
	}
	return nil
}

func TestUpdateReservationControllerSuccess(t *testing.T) {
	// Arrange
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
	createdReservation, err := createReservation(true)
	require.NoError(t, err)
	expectedStatus := 200
	expectedReservation := models.Reservation{
		Id:             "testReservationId",
		Facility:       createdReservation.Facility,
		ParkedVehicle:  createdReservation.ParkedVehicle,
		UpdatePossible: true,
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(reservationInput)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var output models.Reservation
	err = json.Unmarshal(rr.Body.Bytes(), &output)
	require.NoError(t, err)
	assert.Equal(t, expectedReservation, output)

	err = deleteReservation()
	require.NoError(t, err)
}
func TestUpdateReservationControllerNotAllow(t *testing.T) {
	// Arrange
	_, err := createReservation(false)
	require.NoError(t, err)
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

	expectedStatus := 422
	expectedError := lib.ErrorResponse{Code: 422, Title: "unprocessable input", Description: "this reservation can not be update"}

	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(reservationInput)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = deleteReservation()
	require.NoError(t, err)
}

func TestUpdateReservationControllerFailOnGetDynamoDB(t *testing.T) {
	// Arrange
	defer func() {
		lib.AWS_REGION = "us-west-2"
	}()
	lib.AWS_REGION = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestUpdateReservationControllerFailOnNotFound(t *testing.T) {
	// Arrange
	expectedStatus := 404
	expectedError := lib.ErrorResponse{Code: 404, Title: "Not found", Description: "reservation not found"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{
		Id: "01255658855",
	})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestUpdateReservationControllerFailOnCheckUpdateInternalErr(t *testing.T) {
	// Arrange
	defer func() {
		lib.RESERVATION_TABLE_NAME = "ReservationSpot"
	}()
	lib.RESERVATION_TABLE_NAME = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestUpdateReservationControllerFailOnDecodeBody(t *testing.T) {
	// Arrange
	_, err := createReservation(true)
	require.NoError(t, err)
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "please enter correct body request"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal("")
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = deleteReservation()
	require.NoError(t, err)
}

func TestUpdateReservationControllerFailOnIdConflict(t *testing.T) {
	// Arrange
	_, err := createReservation(true)
	require.NoError(t, err)
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "url reservation_id must be equal to body reservation_id"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = deleteReservation()
	require.NoError(t, err)
}

func TestUpdateReservationControllerFailOnValidate(t *testing.T) {
	// Arrange
	_, err := createReservation(true)
	require.NoError(t, err)
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "all fields should be send"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{
		Id: "testReservationId",
	})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = deleteReservation()
	require.NoError(t, err)
}
func TestUpdateReservationControllerFailOnNotFoundFacilityId(t *testing.T) {
	// Arrange
	_, err := createReservation(true)
	require.NoError(t, err)
	expectedStatus := 404
	expectedError := lib.ErrorResponse{Code: 404, Title: "Not found", Description: "facility_id or vehicle_id not found"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{
		Id:            "testReservationId",
		Facility:      "fdfadfas",
		ParkedVehicle: "dfdfdf",
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = deleteReservation()
	require.NoError(t, err)
}

func TestUpdateReservationControllerFailOnFetchReservationInfoErr(t *testing.T) {
	// Arrange
	_, err := createReservation(true)
	require.NoError(t, err)
	defer func() {
		lib.VEHICLE_TABLE_NAME = "VehicleSpot"
		err = deleteReservation()
		require.NoError(t, err)
	}()
	lib.VEHICLE_TABLE_NAME = ""

	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation/{reservation_id}/update", UpdateReservationController).Methods("PUT")
	marshal, err := json.Marshal(models.InputReservation{
		Id:            "testReservationId",
		Facility:      "fdfadfas",
		ParkedVehicle: "dfdfdf",
		Quote: models.Quote{
			Ends:   "10",
			Starts: "5",
		},
	})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPut, "/v2/reservation/testReservationId/update", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

}
