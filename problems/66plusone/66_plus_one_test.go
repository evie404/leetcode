package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},
		{
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.digits), func(t *testing.T) {
			assert.Equal(t, tt.want, plusOne(tt.digits))
		})
	}
}
