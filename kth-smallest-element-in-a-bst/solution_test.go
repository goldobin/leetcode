package kth_smallest_element_in_a_bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "case 0.1 nil node",
			want: -1,
		},
		{
			name: "case 0.2 one node k=1",
			root: &TreeNode{Val: 10},
			k:    1,
			want: 10,
		},
		{
			name: "case 0.3 one node k=2",
			root: &TreeNode{Val: 20},
			k:    2,
			want: -1,
		},
		{
			name: "case 0.4.1 two nodes (left) k=1",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}},
			k:    1,
			want: 10,
		},
		{
			name: "case 0.4.2 two nodes (left) k=2",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}},
			k:    2,
			want: 20,
		},
		{
			name: "case 0.4.3 two nodes (left) k=3",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}},
			k:    3,
			want: -1,
		},
		{
			name: "case 0.5.1 two nodes (right) k=1",
			root: &TreeNode{Val: 20, Right: &TreeNode{Val: 30}},
			k:    1,
			want: 20,
		},
		{
			name: "case 0.5.2 two nodes (right) k=2",
			root: &TreeNode{Val: 20, Right: &TreeNode{Val: 30}},
			k:    2,
			want: 30,
		},
		{
			name: "case 0.5.3 two nodes (right) k=3",
			root: &TreeNode{Val: 20, Right: &TreeNode{Val: 30}},
			k:    3,
			want: -1,
		},
		{
			name: "case 0.6.1 three nodes k=1",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30}},
			k:    1,
			want: 10,
		},
		{
			name: "case 0.6.2 three nodes k=2",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30}},
			k:    2,
			want: 20,
		},
		{
			name: "case 0.6.2 three nodes k=3",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30}},
			k:    3,
			want: 30,
		},
		{
			name: "case 0.6.3 three nodes k=4",
			root: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30}},
			k:    4,
			want: -1,
		},
		{
			name: "case 1.1.1 leetcode 4 nodes k=1",
			root: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 20},
				},
				Right: &TreeNode{Val: 40},
			},
			k:    1,
			want: 10,
		},
		{
			name: "case 1.1.2 leetcode 4 nodes k=2",
			root: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 20},
				},
				Right: &TreeNode{Val: 40},
			},
			k:    2,
			want: 20,
		},
		{
			name: "case 1.1.2 leetcode 4 nodes k=3",
			root: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 20},
				},
				Right: &TreeNode{Val: 40},
			},
			k:    3,
			want: 30,
		},
		{
			name: "case 1.1.2 leetcode 4 nodes k=4",
			root: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 20},
				},
				Right: &TreeNode{Val: 40},
			},
			k:    4,
			want: 40,
		},
		{
			name: "case 1.1.2 leetcode 4 nodes k=5",
			root: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 20},
				},
				Right: &TreeNode{Val: 40},
			},
			k:    5,
			want: -1,
		},
		{
			name: "case 1.2.1 leetcode 6 nodes k=1",
			root: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:  20,
						Left: &TreeNode{Val: 10},
					},
					Right: &TreeNode{Val: 40},
				},
				Right: &TreeNode{
					Val: 60,
				},
			},
			k:    1,
			want: 10,
		},
		{
			name: "case 1.2.2 leetcode 6 nodes k=3",
			root: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:  20,
						Left: &TreeNode{Val: 10},
					},
					Right: &TreeNode{Val: 40},
				},
				Right: &TreeNode{
					Val: 60,
				},
			},
			k:    3,
			want: 30,
		},
		{
			name: "case 1.2.3 leetcode 6 nodes k=6",
			root: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:  20,
						Left: &TreeNode{Val: 10},
					},
					Right: &TreeNode{Val: 40},
				},
				Right: &TreeNode{
					Val: 60,
				},
			},
			k:    6,
			want: 60,
		},
		{
			name: "case 1.2.4 leetcode 6 nodes k=7",
			root: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:  20,
						Left: &TreeNode{Val: 10},
					},
					Right: &TreeNode{Val: 40},
				},
				Right: &TreeNode{
					Val: 60,
				},
			},
			k:    7,
			want: -1,
		},
		{
			name: "case 1.3.1 6 nodes right k=1",
			root: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 50,
					Left: &TreeNode{
						Val:  40,
						Left: &TreeNode{Val: 30},
					},
					Right: &TreeNode{Val: 60},
				},
			},
			k:    1,
			want: 10,
		},
		{
			name: "case 1.3.2 6 nodes right k=2",
			root: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 50,
					Left: &TreeNode{
						Val:  40,
						Left: &TreeNode{Val: 30},
					},
					Right: &TreeNode{Val: 60},
				},
			},
			k:    2,
			want: 20,
		},
		{
			name: "case 1.3.3 6 nodes right k=6",
			root: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 50,
					Left: &TreeNode{
						Val:  40,
						Left: &TreeNode{Val: 30},
					},
					Right: &TreeNode{Val: 60},
				},
			},
			k:    6,
			want: 60,
		},
		{
			name: "case 1.3.4 6 nodes right k=7",
			root: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 50,
					Left: &TreeNode{
						Val:  40,
						Left: &TreeNode{Val: 30},
					},
					Right: &TreeNode{Val: 60},
				},
			},
			k:    7,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallest(tt.root, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
