package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			name: "empty array",
			nums: []int{},
			want: nil,
		},
		{
			name: "single element",
			nums: []int{1},
			want: &TreeNode{Val: 1},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 4},
			},
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{1, 3},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
