package lib

import "errors"

var (
	ErrDatabaseConnection = errors.New("couldn't connect to database")
	ErrNotFound           = errors.New("not found")
	ErrAlreadyExists      = errors.New("already exists")
	ErrForbidden          = errors.New("forbidden")
	ErrInternal           = errors.New("internal error")
)

type ErrorResponse struct {
	Code        int    `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewErrorResponse(code int, title, description string) *ErrorResponse {
	err := new(ErrorResponse)
	err.Code = code
	err.Title = title
	err.Description = description
	return err
}
