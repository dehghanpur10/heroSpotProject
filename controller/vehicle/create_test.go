package vehicle

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
	"spotHeroProject/service/vehicleService"
	"testing"
)

func TestCreateVehicleController(t *testing.T) {
	vehicle := models.Vehicle{
		Id: "123456",
		Description: models.VehicleDescription{
			Name:  "vehicle1",
			Year:  "5",
			Model: "55",
		},
	}
	tests := []struct {
		name   string
		input  models.Vehicle
		status int
		output interface{}
	}{
		{name: "ok", input: vehicle, status: 201, output: vehicle},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			router := mux.NewRouter()
			router.HandleFunc("/v2/vehicle", CreateVehicleController).Methods("POST")
			marshal, err := json.Marshal(test.input)
			require.NoError(t, err)
			req, _ := http.NewRequest(http.MethodPost, "/v2/vehicle", bytes.NewBuffer(marshal))
			rr := httptest.NewRecorder()
			// Act
			router.ServeHTTP(rr, req)
			//Assert
			assert.Equal(t, test.status, rr.Code)
			switch rr.Code {
			case 201:
				var vehicle models.Vehicle
				err = json.Unmarshal(rr.Body.Bytes(), &vehicle)
				require.NoError(t, err)
				assert.Equal(t, test.output.(models.Vehicle), vehicle)
			default:
				var errorResponse lib.ErrorResponse
				err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				require.NoError(t, err)
				assert.Equal(t, test.output.(lib.ErrorResponse), errorResponse)
			}

		})
	}
	db, err := lib.GetDynamoDB()
	require.NoError(t, err)
	service := vehicleService.New(db)
	err = service.DeleteVehicle(vehicle.Id)
	require.NoError(t, err)
}
