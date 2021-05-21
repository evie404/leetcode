package week2

import (
	"math"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	out, err := superpalindromesInRangeErr(left, right)
	if err != nil {
		panic(err)
	}

	return out
}

func superpalindromesInRangeErr(left string, right string) (int, error) {
	leftI, err := strconv.ParseInt(left, 10, 64)
	if err != nil {
		return 0, err
	}

	rightI, err := strconv.ParseInt(right, 10, 64)
	if err != nil {
		return 0, err
	}

	leftSqrt := int64(math.Ceil(math.Sqrt(float64(leftI))))
	rightSqrt := int64(math.Floor(math.Sqrt(float64(rightI))))

	count := 0

	for i := leftSqrt; i <= rightSqrt; i++ {
		if isPalindrome(i*i) && isPalindrome(i) {
			count++
		}
	}

	return count, nil
}

func isPalindrome(x int64) bool {
	if x < 0 {
		return false
	}

	return isPalindromeStr(strconv.FormatInt(x, 10))
}

func isPalindromeStr(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
