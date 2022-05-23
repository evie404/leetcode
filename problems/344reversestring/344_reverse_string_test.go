package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			nil,
			nil,
		},
		{
			[]byte{1},
			[]byte{1},
		},
		{
			[]byte{1, 2},
			[]byte{2, 1},
		},
		{
			[]byte{1, 2, 3},
			[]byte{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.s), func(t *testing.T) {
			reverseString(tt.s)
			assert.Equal(t, tt.want, tt.s)
		})
	}
}
