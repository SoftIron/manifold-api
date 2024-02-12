package sifi

import (
	"fmt"
	"net/http"
	"strings"
)

// ResponseError is the response body for all error responses.
type ResponseError struct {
	Code   int      `json:"code"`
	Text   string   `json:"error"`
	Detail []string `json:"detail"`
}

// NewErrorResponse returns a new ErrorResponse.
func NewErrorResponse(code int, err error, detail ...string) *ResponseError {
	return &ResponseError{
		Code:   code,
		Text:   err.Error(),
		Detail: detail,
	}
}

// Error implements the error interface for e.
func (e ResponseError) Error() string {
	s := fmt.Sprintf("%s (%d) - %s", strings.ToLower(http.StatusText(e.Code)), e.Code, e.Text)
	for _, d := range e.Detail {
		s += "\n\t" + d
	}

	return s
}

// StatusCode returns the status code for e.
func (e ResponseError) StatusCode() int {
	return e.Code
}
