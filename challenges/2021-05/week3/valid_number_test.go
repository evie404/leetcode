package week3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	trueCases := []string{
		"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
	}
	falseCases := []string{
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", ".", "-", "+", "", " ",
	}

	for _, trueCase := range trueCases {
		t.Run(trueCase+" is a number", func(t *testing.T) {
			got := isNumber(trueCase)
			assert.True(t, got)
		})
	}

	for _, falseCase := range falseCases {
		t.Run(falseCase+" is not a number", func(t *testing.T) {
			got := isNumber(falseCase)
			assert.False(t, got)
		})
	}
}
