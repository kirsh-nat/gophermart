package order

import (
	"fmt"
	"strings"
)

type UserNumberExistsError struct {
	level string
	Err   error
}
type NumberExists struct {
	level string
	Err   error
}

type NumberFormatError struct {
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

func (le *NumberExists) Error() string {
	return fmt.Sprintf("[%s] %v", le.level, le.Err)
}

func NewNumberExists(label string, err error) error {
	return &NumberExists{
		level: strings.ToUpper(label),
		Err:   err,
	}
}

func (le *NumberFormatError) Error() string {
	return fmt.Sprintf("[%s] %v", le.level, le.Err)
}

func NewNumberFormatError(label string, err error) error {
	return &NumberFormatError{
		level: strings.ToUpper(label),
		Err:   err,
	}
}
