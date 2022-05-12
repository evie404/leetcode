package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.want, majorityElement(tt.nums))
		})
	}
}
