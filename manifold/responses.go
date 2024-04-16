package manifold

import (
	"fmt"
	"net/http"
	"strings"
)

// ResponseError is the response body for all error responses.
type ResponseError struct {
	Code    int      `json:"code"`
	Text    string   `json:"error"`
	Details []string `json:"details"`
}

// InfoResponse is the response body for GET /.
type InfoResponse struct {
	ProductName    string `json:"product_name"`
	ProductRelease string `json:"product_release"`
}

// NewErrorResponse returns a new ErrorResponse.
func NewErrorResponse(code int, err error, detail ...string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Text:    err.Error(),
		Details: detail,
	}
}

// Error implements the error interface for e.
func (e ResponseError) Error() string {
	s := fmt.Sprintf("%s (%d)", strings.ToLower(http.StatusText(e.Code)), e.Code)
	if e.Text != "" {
		s += " - " + e.Text
	}
	for _, d := range e.Details {
		s += "\n\t" + d
	}

	return s
}

// StatusCode returns the status code for e.
func (e ResponseError) StatusCode() int {
	return e.Code
}
