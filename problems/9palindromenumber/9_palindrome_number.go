package palindromenumber

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return isPalindromeStr(strconv.Itoa(x))
}

func isPalindromeStr(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
