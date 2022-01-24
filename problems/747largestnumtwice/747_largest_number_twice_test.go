package largestnumbertwice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"",
			[]int{3, 6, 1, 0},
			1,
		},
		{
			"",
			[]int{1, 2, 3, 4},
			-1,
		},
		{
			"",
			[]int{1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dominantIndex(tt.nums))
		})
	}
}
