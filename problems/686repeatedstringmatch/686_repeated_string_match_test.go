package repeatedstringmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repeatedStringMatch(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{
			"aaaaaaaaaaaaaaaaaaaaaab",
			"ba",
			2,
		},
		{
			"a",
			"a",
			1,
		},
		{
			"a",
			"aa",
			2,
		},
		{
			"aa",
			"a",
			1,
		},
		{
			"a",
			"bbb",
			-1,
		},
		{
			"abcd",
			"babcdbabcdb",
			-1,
		},
		{
			"abcd",
			"abcdabcdabcd",
			3,
		},
		{
			"abcd",
			"cdabcdab",
			3,
		},
		{
			"bb",
			"bbbbbbb",
			4,
		},
		{
			"abcd",
			"bc",
			1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, repeatedStringMatch(tt.a, tt.b))
		})
	}
}
