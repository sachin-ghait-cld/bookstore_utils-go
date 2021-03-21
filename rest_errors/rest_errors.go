package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RestErr generic error
type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

func (r restErr) Message() string {
	return r.message
}
func (r restErr) Status() int {
	return r.status
}
func (r restErr) Error() string {
	return fmt.Sprintf("message: %s - status %d - error: %s - causes: [%v]", r.message, r.status, r.error, r.causes)
}
func (r restErr) Causes() []interface{} {
	return r.causes
}

// NewRestError return NewRestError
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

// NewRestErrorFromBytes return NewRestErrorFromBytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// NewBadRequestError return bad request
func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

// NewNotFoundError return not found
func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

// NewUnauthorizedError return not found
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

// NewInternalServerError internal_server_error
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
