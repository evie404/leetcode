package allintwobst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
				Val: 2,
			},
			[]int{1, 2, 4},
		},
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			},
			[]int{0, 1, 3},
		},
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Val: 8,
			},
			[]int{1, 8},
		},
		{
			&TreeNode{
				Right: &TreeNode{
					Val: 8,
				},
				Val: 1,
			},
			[]int{1, 8},
		},
		{
			&TreeNode{
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 8,
					},
					Left: &TreeNode{
						Val: 3,
					},
				},
				Val: 1,
			},
			[]int{1, 3, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.root), func(t *testing.T) {
			assert.Equal(t, tt.want, inOrder(tt.root))
		})
	}
}

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 4},
			[]int{3, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v + %+v", tt.a, tt.b), func(t *testing.T) {
			assert.Equal(t, tt.want, mergeSort(tt.a, tt.b))
		})
	}
}

func Test_getAllElements(t *testing.T) {
	tests := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  []int
	}{
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
				Val: 2,
			},
			&TreeNode{
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			},
			[]int{0, 1, 1, 2, 3, 4},
		},
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			},
			&TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
				Val: 2,
			},
			[]int{0, 1, 1, 2, 3, 4},
		},
		{
			&TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Val: 8,
			},
			&TreeNode{
				Right: &TreeNode{
					Val: 8,
				},
				Val: 1,
			},
			[]int{1, 1, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v + %+v", tt.root1, tt.root2), func(t *testing.T) {
			assert.Equal(t, tt.want, getAllElements(tt.root1, tt.root2))
		})
	}
}
