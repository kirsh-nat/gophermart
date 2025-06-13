package draftservices

import (
	"fmt"
	"strings"
)

type UserNumberExistsError struct {
	level string
	Err   error
}
type UserNotAuthorizedError struct {
	level string
	Err   error
}

type PaymentRequiredError struct {
	level string
	Err   error
}

func (le *UserNumberExistsError) Error() string {
	return fmt.Sprintf("[%s] %v", le.level, le.Err)
}

func NewUserNumberExistsError(label string, err error) error {
	return &UserNumberExistsError{
		level: strings.ToUpper(label),
		Err:   err,
	}
}

func (le *UserNotAuthorizedError) Error() string {
	return fmt.Sprintf("[%s] %v", le.level, le.Err)
}

func NewUserNotAuthorizedError(label string, err error) error {
	return &UserNotAuthorizedError{
		level: strings.ToUpper(label),
		Err:   err,
	}
}

func (le *PaymentRequiredError) Error() string {
	return fmt.Sprintf("[%s] %v", le.level, le.Err)
}

func NewPaymentRequiredError(label string, err error) error {
	return &PaymentRequiredError{
		level: strings.ToUpper(label),
		Err:   err,
	}
}
