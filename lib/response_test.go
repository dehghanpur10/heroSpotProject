package lib

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)





func TestJsonResponseHeaders(t *testing.T) {
	// Arrange
	expectedAllowOrigin := "*"
	expectedContentType := "application/json"
	rr := httptest.NewRecorder()
	// Act
	jsonResponseHeaders(rr)
	// Assert
	assert.Equal(t, expectedAllowOrigin, rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, expectedContentType, rr.Header().Get("Content-Type"))
}

func TestHttpSuccessResponse(t *testing.T) {
	// Arrange
	payload := []byte(`{
					"name":"mohammad",
					"age":22	
				}`)
	rr := httptest.NewRecorder()
	// Act
	HttpSuccessResponse(rr, 200, payload)
	//Assert
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, payload, rr.Body.Bytes())
}

func TestHttpErrorResponse(t *testing.T) {
	// Arrange
	expectedStatusCode := 400
	expectedTitleError := "bad request"
	expectedDescription := "user input isn't correct"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()
	// Act
	HttpErrorResponse(rr, expectedStatusCode, expectedTitleError, expectedDescription)
	// Assert
	assert.Equal(t, expectedResponseBody, RemoveLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError400(t *testing.T) {
	// Arrange
	expectedStatusCode := 400
	expectedTitleError := "Bad request"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()
	// Act
	HttpError400(rr, expectedDescription)
	// Assert
	assert.Equal(t, expectedResponseBody, RemoveLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError404(t *testing.T) {
	// Arrange
	expectedStatusCode := 404
	expectedTitleError := "Not found"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()
	// Act
	HttpError404(rr, expectedDescription)
	// Assert
	assert.Equal(t, expectedResponseBody, RemoveLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}
func TestHttpError422(t *testing.T) {
	// Arrange
	expectedStatusCode := 422
	expectedTitleError := "unprocessable input"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()
	// Act
	HttpError422(rr, expectedDescription)
	// Assert
	assert.Equal(t, expectedResponseBody, RemoveLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError500(t *testing.T) {
	// Arrange
	expectedStatusCode := 500
	expectedTitleError := "Internal error"
	expectedDescription := "Internal server error."
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()
	// Act
	HttpError500(rr)
	// Assert
	assert.Equal(t, expectedResponseBody, RemoveLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}
