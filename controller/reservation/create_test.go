package reservation

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/facilityService"
	"spotHeroProject/service/reservationService"
	"spotHeroProject/service/vehicleService"
	"testing"
)

func TestCreateReservationControllerSuccess(t *testing.T) {
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
	err := os.Setenv("AWS_REGION", "us-west-2")
	require.NoError(t, err)
	db, err := lib.GetDynamoDB()
	require.NoError(t, err)
	serviceOfVehicle := vehicleService.New(db)
	err = serviceOfVehicle.CreateVehicle(vehicle)
	require.NoError(t, err)
	serviceOfFacility := facilityService.New(db)
	err = serviceOfFacility.CreateFacilityService(facility)
	require.NoError(t, err)

	expectedStatus := 201
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
	router.HandleFunc("/v2/reservation", CreateReservationController).Methods("POST")
	marshal, err := json.Marshal(reservationInput)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/reservation", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var reservation models.Reservation
	err = json.Unmarshal(rr.Body.Bytes(), &reservation)
	require.NoError(t, err)
	assert.Equal(t, expectedReservation, reservation)

	err = serviceOfFacility.DeleteFacilityService(facility.Id)
	require.NoError(t, err)
	err = serviceOfVehicle.DeleteVehicle(vehicle.Id)
	require.NoError(t, err)
	serviceOfReservation := reservationService.New(db)
	err = serviceOfReservation.DeleteReservationService(reservationInput.Id)
	require.NoError(t, err)
	err = os.Unsetenv("AWS_REGION")
	require.NoError(t, err)
}
func TestCreateReservationControllerFailOnDecode(t *testing.T) {
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "please enter correct body request"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", CreateReservationController).Methods("POST")
	marshal, err := json.Marshal("")
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/reservation", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestCreateReservationControllerFailOnValidate(t *testing.T) {
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "all fields should be send"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", CreateReservationController).Methods("POST")
	marshal, err := json.Marshal(models.Reservation{})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/reservation", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}
func TestCreateReservationControllerFailOnGetDynamodb(t *testing.T) {
	// Arrange
	reservation := models.InputReservation{
		Id:             "1",
		ParkedVehicle:  "1",
		Facility:       "1",
		UpdatePossible: false,
		Quote: models.Quote{
			Ends:   "1",
			Starts: "0",
		},
	}
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", CreateReservationController).Methods("POST")
	marshal, err := json.Marshal(reservation)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/reservation", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)
}

func TestCreateReservationControllerFailOnNotFound(t *testing.T) {
	err := os.Setenv("AWS_REGION", "us-west-2")
	require.NoError(t, err)
	reservation := models.InputReservation{
		Id:             "1",
		ParkedVehicle:  "1000520",
		Facility:       "1",
		UpdatePossible: false,
		Quote: models.Quote{
			Ends:   "1",
			Starts: "0",
		},
	}
	expectedStatus := 404
	expectedError := lib.ErrorResponse{Code: 404, Title: "Not found", Description: "facility_id or vehicle_id not found"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", CreateReservationController).Methods("POST")
	marshal, err := json.Marshal(reservation)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/reservation", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var errorResponse lib.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, expectedError, errorResponse)

	err = os.Unsetenv("AWS_REGION")
	require.NoError(t, err)
}