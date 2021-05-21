package week1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			"eg1",
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			"eg2",
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"eg3",
			[]int{3, 1, 2, 10, 1},
			[]int{3, 4, 6, 16, 17},
		},
		{
			"empty",
			[]int{},
			[]int{},
		},
		{
			"nil",
			nil,
			[]int{},
		},
		{
			"one number",
			[]int{12},
			[]int{12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runningSum(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
