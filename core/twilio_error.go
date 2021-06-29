package core

import (
	"errors"
	"fmt"
)

var TwilioErrorGeneric = errTwilio()

func errTwilio() error {
	return TwilioError{}
}

// TwilioError provides information about an unsuccessful request.
type TwilioError struct {
	Cause   error
	Status  int
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err TwilioError) Error() string {
	return fmt.Sprintf("Status: %d - Error %d: %s",
		err.Status, err.Code, err.Message)
}

func (err TwilioError) Is(target error) bool {
	_, ok := target.(TwilioError)
	return ok
}

func WrapError(err error, status int, code int, message string) error {
	return TwilioError{
		Cause:   err,
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func WrapErrorStatus(err error, status int, message string) error {
	return WrapError(err, status, 45001, message)
}

func WrapErrorGeneric(err error, message string) error {
	return WrapError(err, 500, 45001, message)
}

func CreateErrorGeneric(message string) error {
	return WrapError(errors.New(""), 500, 45001, message)
}
