package detectcapital

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"empty string",
			args{""},
			true,
		},
		{
			"all caps",
			args{"USA"},
			true,
		},
		{
			"all lowercase",
			args{"leetcode"},
			true,
		},
		{
			"only first character cap",
			args{"Google"},
			true,
		},
		{
			"mixed cap",
			args{"FlaG"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, detectCapitalUse(tt.args.word))
		})
	}
}
