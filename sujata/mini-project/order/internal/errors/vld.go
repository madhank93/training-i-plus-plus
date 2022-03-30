package errors

import (
	"net/http"
	"order/internal/literals"
)

// ServerError defines the error message and related http response code for any
// error return by the microservice.
type ServerError struct {
	ErrorMessage     string
	HttpResponseCode int
}

// Validation errors
var (
	ParametersMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.ParametersMissing,
		HttpResponseCode: http.StatusBadRequest,
	}
)
