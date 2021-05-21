package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			"0",
			0,
			true,
		},
		{
			"121",
			121,
			true,
		},
		{
			"-121",
			-121,
			false,
		},
		{
			"11",
			11,
			true,
		},
		{
			"10",
			10,
			false,
		},
		{
			"-101",
			-101,
			false,
		},
		{
			"1000021",
			1000021,
			false,
		},
		{
			"0",
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
