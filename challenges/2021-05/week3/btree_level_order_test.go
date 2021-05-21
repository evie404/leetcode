package week3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		arg  *TreeNode
		want [][]int
	}{
		{
			"nothing",
			nil,
			[][]int{},
		},
		{
			"just root node",
			&TreeNode{Val: 1},
			[][]int{{1}},
		},
		{
			"one level",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			[][]int{{1}, {2, 3}},
		},
		{
			"two levels",
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 21}, Right: &TreeNode{Val: 22}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 31}}},
			[][]int{{1}, {2, 3}, {21, 22, 31}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
