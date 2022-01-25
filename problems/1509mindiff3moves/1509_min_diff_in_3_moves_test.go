package mindiff3moves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDifference(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{5, 3, 2, 4},
			0,
		},
		{
			[]int{1, 5, 0, 10, 14},
			1,
		},
		{
			[]int{0, 1, 1, 4, 6, 6, 6},
			2,
		},
		{
			[]int{0, 0, 0, 0, 0, 1, 4, 4, 4, 4, 4, 4, 6, 6, 6, 6},
			6,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.want, minDifference(tt.nums))
		})
	}
}
