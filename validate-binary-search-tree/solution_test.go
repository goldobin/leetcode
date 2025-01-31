package validate_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "case 0.1",
			want: true,
		},
		{
			name: "case 0.2",
			root: &TreeNode{
				Val: 1,
			},
			want: true,
		},
		{
			name: "case 1.1 leetcode",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			name: "case 1.2 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: false,
		},
		{
			name: "case 1.3",
			root: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: true,
		},
		{
			name: "case 1.4",
			root: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
