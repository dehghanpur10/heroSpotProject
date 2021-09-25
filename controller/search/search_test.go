package search

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"spotHeroProject/lib"
	"testing"
)

func TestSearchFacilityControllerSuccess(t *testing.T) {
	// Arrange
	err := os.Setenv("AWS_REGION", "us-west-2")
	require.NoError(t, err)
	expectedStatus := 200
	router := mux.NewRouter()
	router.HandleFunc("/v2/search", SearchFacilityController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/search", nil)
	rr := httptest.NewRecorder()
	// Act
	router.ServeHTTP(rr, req)
	//Assert
	assert.Equal(t, expectedStatus, rr.Code)
	err = os.Unsetenv("AWS_REGION")
	require.NoError(t, err)
}

func TestSearchFacilityControllerFailOnGetDynamodb(t *testing.T) {
	// Arrange
	expectedStatus := 500
	expectedError := lib.ErrorResponse{Code: 500, Title: "Internal error", Description: "Internal server error."}
	router := mux.NewRouter()
	router.HandleFunc("/v2/search", SearchFacilityController).Methods("GET")
	req, _ := http.NewRequest(http.MethodGet, "/v2/search", nil)
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

func TestSearchFacilityControllerFailOnQueryParsing(t *testing.T) {
	err := os.Setenv("AWS_REGION", "us-west-2")
	require.NoError(t, err)
	tests := []struct {
		name          string
		url           string
		expectedError string
	}{
		{name: "latitude", url: "/v2/search?lat=5d&lon=1", expectedError: "lat query should be number type"},
		{name: "longitude", url: "/v2/search?lat=5&lon=1d", expectedError: "lon query should be number type"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			expectedStatus := 400
			expectedError := lib.ErrorResponse{Code: 400, Title: "Bad request", Description: test.expectedError}
			router := mux.NewRouter()
			router.HandleFunc("/v2/search", SearchFacilityController).Methods("GET")
			req, _ := http.NewRequest(http.MethodGet, test.url, nil)
			rr := httptest.NewRecorder()
			// Act
			router.ServeHTTP(rr, req)
			//Assert
			assert.Equal(t, expectedStatus, rr.Code)
			var errorResponse lib.ErrorResponse
			err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
			require.NoError(t, err)
			assert.Equal(t, expectedError, errorResponse)

		})
	}
	err = os.Unsetenv("AWS_REGION")
	require.NoError(t, err)

}
