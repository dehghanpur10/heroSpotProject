package reservation

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"spotHeroProject/lib"
	"testing"
)

func TestGetAllReservationControllerSuccess(t *testing.T) {
	// Arrange
	expectedStatus := 200
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", GetAllReservationController).Methods("GET")

	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
}

func TestGetAllReservationControllerFailOnGetDynamodb(t *testing.T) {
	// Arrange
	defer func() {
		lib.AWS_REGION = "us-west-2"
	}()
	lib.AWS_REGION = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", GetAllReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation", nil)
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

func TestGetAllReservationControllerFailOnGetAllReservationService(t *testing.T) {
	// Arrange
	defer func() {
		lib.RESERVATION_TABLE_NAME = "ReservationSpot"
	}()
	lib.RESERVATION_TABLE_NAME = ""
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/reservation", GetAllReservationController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/reservation", nil)
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
