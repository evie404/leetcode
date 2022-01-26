package totalhammingdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingDistance(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			2, 1, 2,
		},
		{
			6, 3, 2,
		},
		{
			12, 3, 4,
		},
		{
			24, 3, 4,
		},
		{
			4, 1, 2,
		},
		{
			8, 1, 2,
		},
		{
			4, 14, 2,
		},
		{
			4, 2, 2,
		},
		{
			2, 14, 2,
		},
		{
			2, 2, 0,
		},
		{
			3, 3, 0,
		},
		{
			6, 1, 3,
		},
		{
			111, 1, 5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.a, tt.b), func(t *testing.T) {
			assert.Equal(t, tt.want, hammingDistance(tt.a, tt.b))
		})

		t.Run(fmt.Sprintf("%v, %v", tt.b, tt.a), func(t *testing.T) {
			assert.Equal(t, tt.want, hammingDistance(tt.b, tt.a))
		})
	}
}

func Test_totalHammingDistance(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{4, 14, 2},
			6,
		},
		{
			[]int{4, 14, 4},
			4,
		},
		{
			[]int{6, 1},
			3,
		},
		{
			[]int{6, 1, 8},
			8,
		},
		{
			[]int{6, 1, 8, 6},
			14,
		},
		{
			[]int{6, 1, 8, 6, 8},
			22,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.want, totalHammingDistance(tt.nums))
		})
	}
}
