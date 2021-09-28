package httpserver

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

func TestNotFoundHandler(t *testing.T) {
	// Arrange
	expectedResponse := lib.NewErrorResponse(404, "Not found", "Requested resource doesn't exist. Please check your path.")
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFoundController)
	req, err := http.NewRequest(http.MethodGet, "/devices/", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	// Act
	router.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, 404, rr.Code)
	body, err := json.Marshal(expectedResponse)
	require.NoError(t, err)
	assert.Equal(t, body, lib.RemoveLineBreakTrailing(rr.Body.Bytes()))
}
