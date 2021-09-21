package lib

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO (@ryan.bm): Its better to remove extra lines (like line 17 and 19)
// TODO (@ryan.bm): format document before pushing to git (use gofmt or prettire in vscode)
func removeLineBreakTrailing(data []byte) []byte {

	return []byte(strings.TrimSuffix(string(data), "\n"))

}

// TODO (@ryan.bm): Its better to put the assertion inside the test function, delete headerTest()
func headerTest(w http.ResponseWriter, t *testing.T) {
	expectedAllowOrigin := "*"
	expectedContentType := "application/json"
	assert.Equal(t, expectedAllowOrigin, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, expectedContentType, w.Header().Get("Content-Type"))
}

func TestJsonResponseHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonResponseHeaders(rr)
	headerTest(rr, t)
}

/*
	TODO (@ryan.bm): Its better to separate the different parts of the test with these comments:
	// Arrange
	...
	// Act
	...
	//Assert
	...
*/
func TestHttpSuccessResponse(t *testing.T) {
	payload := []byte(`{
					"name":"mohammad",
					"age":22	
				}`)
	rr := httptest.NewRecorder()

	HttpSuccessResponse(rr, 200, payload)
	headerTest(rr, t)

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, payload, rr.Body.Bytes())
}

func TestHttpErrorResponse(t *testing.T) {
	expectedStatusCode := 400
	expectedTitleError := "bad request"
	expectedDescription := "user input isn't correct"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	HttpErrorResponse(rr, expectedStatusCode, expectedTitleError, expectedDescription)

	assert.Equal(t, expectedResponseBody, removeLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError400(t *testing.T) {
	expectedStatusCode := 400
	expectedTitleError := "Bad request"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	HttpError400(rr, expectedDescription)

	assert.Equal(t, expectedResponseBody, removeLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError404(t *testing.T) {
	expectedStatusCode := 404
	expectedTitleError := "Not found"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	HttpError404(rr, expectedDescription)

	assert.Equal(t, expectedResponseBody, removeLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}
func TestHttpError422(t *testing.T) {
	expectedStatusCode := 422
	expectedTitleError := "unprocessable input"
	expectedDescription := "description"
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	HttpError422(rr, expectedDescription)

	assert.Equal(t, expectedResponseBody, removeLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}

func TestHttpError500(t *testing.T) {
	expectedStatusCode := 500
	expectedTitleError := "Internal error"
	expectedDescription := "Internal server error."
	expectedError := NewErrorResponse(expectedStatusCode, expectedTitleError, expectedDescription)
	expectedResponseBody, err := json.Marshal(expectedError)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	HttpError500(rr)

	assert.Equal(t, expectedResponseBody, removeLineBreakTrailing(rr.Body.Bytes()))
	assert.Equal(t, expectedStatusCode, rr.Code)
}
