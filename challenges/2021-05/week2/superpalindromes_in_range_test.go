package week2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_superpalindromesInRange(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ex1",
			args{
				"4",
				"1000",
			},
			4,
		},
		{
			"ex2",
			args{
				"1",
				"2",
			},
			1,
		},
		{
			"long ex",
			args{
				"38455498359",
				"999999999999999999",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := superpalindromesInRange(tt.args.left, tt.args.right)
			assert.Equal(t, tt.want, got)
		})
	}
}
