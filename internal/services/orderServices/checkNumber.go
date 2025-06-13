package orderservices

import (
	"errors"
	"unicode"
)

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func luhnCheck(number string) bool {
	sum := 0
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}

func CheckNumber(input string) error {
	if !isNumeric(input) {
		return NewNumberFormatError("CheckNumber", errors.New("not a number"))
	}
	if !luhnCheck(input) {
		return NewNumberFormatError("CheckNumber", errors.New("not a number"))
	}
	return nil
}
