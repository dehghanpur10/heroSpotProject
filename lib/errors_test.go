package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	code        = 400
	title       = "bad request"
	description = "user input isn't correct"
	expected    = ErrorResponse{
		400, "bad request", "user input isn't correct",
	}
)

func TestNewErrorResponse(t *testing.T) {
	err := NewErrorResponse(code, title, description)
	assert.Equal(t, &expected, err)
}
