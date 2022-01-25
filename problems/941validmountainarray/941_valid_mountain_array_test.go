package validmountainarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{
			[]int{},
			false,
		},
		{
			[]int{2},
			false,
		},
		{
			[]int{2, 1},
			false,
		},
		{
			[]int{1, 2, 1},
			true,
		},
		{
			[]int{3, 5, 5},
			false,
		},
		{
			[]int{0, 3, 2, 1},
			true,
		},
		{
			[]int{0, 2, 3, 4, 5, 2, 1, 0},
			true,
		},
		{
			[]int{0, 2, 3, 3, 5, 2, 1, 0},
			false,
		},
		{
			[]int{0, 1, 2, 3},
			false,
		},
		{
			[]int{3, 2, 1, 0},
			false,
		},
		{
			[]int{0, 1, 2, 3, 2, 1, 0},
			true,
		},
		{
			[]int{0, 1, 2, 3, 2, 1, 2, 3},
			false,
		},
		{
			[]int{0, 1, 2, 3, 2, 1, 2, 3, 2, 1, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.arr), func(t *testing.T) {
			assert.Equal(t, tt.want, validMountainArray(tt.arr))
		})
	}
}
