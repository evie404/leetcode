package largestmagicsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestMagicSquare(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"1",
			[][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}},
			3,
		},
		{
			"2",
			[][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}},
			2,
		},
		{
			"3",
			[][]int{{3, 3, 3, 3}, {3, 3, 3, 3}, {3, 3, 3, 3}, {3, 3, 3, 3}},
			4,
		},
		{
			"4",
			[][]int{{1}},
			1,
		},
		{
			"5",
			[][]int{{1, 1}, {1, 1}},
			2,
		},
		{
			"6",
			[][]int{{5, 1, 6}, {5, 4, 3}, {2, 7, 3}},
			3,
		},
		{
			"7",
			[][]int{{5, 1, 6}, {5, 4, 1}, {2, 1, 3}},
			1,
		},
		{
			"8",
			[][]int{{3, 3, 3, 3}, {3, 3, 3, 3}, {3, 3, 3, 1}, {3, 3, 1, 3}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestMagicSquare(tt.grid))
		})
	}
}
