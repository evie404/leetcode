package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			listToLinkedList([]int{1, 1, 1}),
			listToLinkedList([]int{1, 1, 1}),
		},
		{
			listToLinkedList([]int{1, 2, 3}),
			listToLinkedList([]int{3, 2, 1}),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(tt.head))
		})
	}
}
