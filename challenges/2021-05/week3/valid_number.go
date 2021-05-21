package week3

import (
	"strings"
	"unicode"
)

func isNumber(s string) bool {
	sLower := strings.ToLower(s)

	if strings.Contains(sLower, "e") {
		parts := strings.Split(sLower, "e")
		if len(parts) != 2 {
			return false
		}

		return isDecimalNumber(parts[0]) && isWholeNumber(parts[1])
	}

	return isDecimalNumber(sLower)
}

func isDecimalNumber(s string) bool {
	if s == "" {
		return false
	}

	hasDecimal := false
	hasDigitsAfterDecimal := false
	hasDigitsBeforeDecimal := false

	hasSign := false
	hasDigitsAfterSign := false

	for i, c := range s {
		if i == 0 && isSign(c) {
			hasSign = true

			continue
		}

		if c == '.' {
			if hasDecimal {
				return false
			}

			hasDecimal = true
			continue
		}

		if !isDigit(c) {
			return false
		}

		if hasSign {
			hasDigitsAfterSign = true
		}

		if hasDecimal {
			hasDigitsAfterDecimal = true
		} else {
			hasDigitsBeforeDecimal = true
		}
	}

	return (!hasSign || hasDigitsAfterSign) && (!hasDecimal || hasDigitsAfterDecimal || hasDigitsBeforeDecimal)
}

func isWholeNumber(s string) bool {
	if s == "" {
		return false
	}

	hasSign := false
	hasDigitsAfterSign := false

	for i, c := range s {
		if i == 0 && isSign(c) {
			hasSign = true

			continue
		}

		if !isDigit(c) {
			return false
		}

		if hasSign {
			hasDigitsAfterSign = true
		}
	}

	return !hasSign || hasDigitsAfterSign
}

func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func isSign(c rune) bool {
	return c == '-' || c == '+'
}
