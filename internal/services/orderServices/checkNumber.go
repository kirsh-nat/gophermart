package orderservices

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/kirsh-nat/gophermart.git/gophermart/internal/models/order"
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

	parts := strings.Fields(input)

	for _, part := range parts {
		if !isNumeric(part) {
			fmt.Print(1111)
			return order.NewNumberFormatError("CheckNumber", fmt.Errorf("'%s' is not a number", part))
		}

		if luhnCheck(part) {
			return order.NewNumberFormatError("CheckNumber", fmt.Errorf("'%s' is not a number", part))
		} else {
			continue
		}
	}

	return nil
}
