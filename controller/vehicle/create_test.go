package vehicle

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
	"spotHeroProject/service/vehicleService"
	"testing"
)

func TestCreateVehicleControllerSuccess(t *testing.T) {
	err := os.Setenv("AWS_REGION", "us-west-2")
	require.NoError(t, err)

	vehicle := models.Vehicle{
		Id: "123456",
		Description: models.VehicleDescription{
			Name:  "vehicle1",
			Year:  "5",
			Model: "55",
		},
	}
	expectedStatus := 201
	router := mux.NewRouter()
	router.HandleFunc("/v2/vehicle", CreateVehicleController).Methods("POST")
	marshal, err := json.Marshal(vehicle)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/vehicle", bytes.NewBuffer(marshal))
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	var output models.Vehicle
	err = json.Unmarshal(rr.Body.Bytes(), &output)
	require.NoError(t, err)
	assert.Equal(t, vehicle, output)

	db, err := lib.GetDynamoDB()
	require.NoError(t, err)
	service := vehicleService.New(db)
	err = service.DeleteVehicle(vehicle.Id)
	require.NoError(t, err)
	err = os.Unsetenv("AWS_REGION")
	require.NoError(t, err)
}

func TestCreateVehicleControllerFailOnDecode(t *testing.T) {
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "please enter correct body request"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/vehicle", CreateVehicleController).Methods("POST")
	marshal, err := json.Marshal("")
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/vehicle", bytes.NewBuffer(marshal))
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
func TestCreateVehicleControllerFailOnValidate(t *testing.T) {
	expectedStatus := 400
	expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: "all fields should be send"}
	router := mux.NewRouter()
	router.HandleFunc("/v2/vehicle", CreateVehicleController).Methods("POST")
	marshal, err := json.Marshal(models.Vehicle{})
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/vehicle", bytes.NewBuffer(marshal))
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

func TestCreateVehicleControllerFailOnGetDynamodb(t *testing.T) {
	vehicle := models.Vehicle{
		Id: "123456",
		Description: models.VehicleDescription{
			Name:  "vehicle1",
			Year:  "5",
			Model: "55",
		},
	}
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/vehicle", CreateVehicleController).Methods("POST")
	marshal, err := json.Marshal(vehicle)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v2/vehicle", bytes.NewBuffer(marshal))
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

