package validpalindrome

import (
	"regexp"
	"strings"
)

var (
	alphaRegexp = regexp.MustCompile(`[^0-9A-Za-z]`)
)

func isPalindrome(s string) bool {
	lowerS := strings.ToLower(s)
	sanatizedS := alphaRegexp.ReplaceAllString(lowerS, "")

	return isPalindromeStr(sanatizedS)
}

func isPalindromeStr(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
