package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			"amanaplanacanalpanama",
			true,
		},
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			"raceacar",
			false,
		},
		{
			"00",
			true,
		},
		{
			"0f",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := isPalindrome(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
