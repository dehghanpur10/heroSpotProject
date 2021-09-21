package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO (@ryan.bm): why this isn't const?
var (
	code        = 400
	title       = "bad request"
	description = "user input isn't correct"
	expected    = ErrorResponse{
		400, "bad request", "user input isn't correct",
	}
)

func TestNewErrorResponse(t *testing.T) {
	/*
		TODO (@ryan.bm): use the variables here like this:
		expectedCode = 400
		expectedTitle = "title"
		...
	*/
	err := NewErrorResponse(code, title, description)
	assert.Equal(t, &expected, err)
}
