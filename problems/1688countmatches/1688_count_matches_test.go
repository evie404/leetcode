package countmatches

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfMatches(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			7,
			6,
		},
		{
			14,
			13,
		},
		{
			3,
			2,
		},
		{
			2,
			1,
		},
		{
			1,
			0,
		},
		{
			0,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Equal(t, tt.want, numberOfMatches(tt.n))
		})
	}
}
