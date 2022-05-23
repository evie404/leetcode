package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			4, 2,
		},
		{
			9, 3,
		},
		{
			10, 3,
		},
		{
			8, 2,
		},
		{
			0, 0,
		},
		{
			1, 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.x), func(t *testing.T) {
			assert.Equal(t, tt.want, mySqrt(tt.x))
		})
	}
}
