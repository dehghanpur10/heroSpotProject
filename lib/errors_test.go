package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorResponse(t *testing.T) {

	code := 400
	title := "bad request"
	description := "user input isn't correct"
	expected := ErrorResponse{
		400, "bad request", "user input isn't correct",
	}

	err := NewErrorResponse(code, title, description)
	assert.Equal(t, &expected, err)
}
