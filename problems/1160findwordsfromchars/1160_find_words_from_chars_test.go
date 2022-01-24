package findwordsfromchars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countCharacters(t *testing.T) {
	tests := []struct {
		words []string
		chars string
		want  int
	}{
		{
			[]string{
				"cat", "bt", "hat", "tree",
			},
			"atach",
			6,
		},
		{
			[]string{
				"hello", "world", "leetcode",
			},
			"welldonehoneyr",
			10,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, countCharacters(tt.words, tt.chars))
		})
	}
}
