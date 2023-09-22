package sifi

// ErrorResponse is the response body for all error responses.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
