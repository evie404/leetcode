package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subtractProductAndSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			234,
			15,
		},
		{
			4421,
			21,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
			assert.Equal(t, tt.want, subtractProductAndSum(tt.n))
		})
	}
}
