package path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      bool
	}{
		{
			name:      "case 0.0 nil tree targetSum 0",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name:      "case 0.1 nil tree targetSum 1",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name:      "case 0.2.2 one node tree",
			root:      &TreeNode{Val: 1},
			targetSum: 0,
			want:      false,
		},
		{
			name:      "case 0.2.3 one node tree",
			root:      &TreeNode{Val: 1},
			targetSum: 2,
			want:      false,
		},
		{
			name: "case 1.1.1 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			want:      true,
		},
		{
			name: "case 1.1.2 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 5,
			want:      false,
		},
		{
			name: "case 1.1.3 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 26,
			want:      true,
		},
		{
			name: "case 1.1.4 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 26,
			want:      true,
		},
		{
			name: "case 1.1.5 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 14,
			want:      false,
		},
		{
			name: "case 1.1.6 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 4,
			want:      false,
		},
		{
			name: "case 1.1.7 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 8,
			want:      false,
		},
		{
			name: "case 1.1.8 leetcode",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: -1,
			want:      false,
		},
		{
			name: "case 1.2.1 negative nodes",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: -7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: -13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: -1,
						},
					},
				},
			},
			targetSum: 14,
			want:      true,
		},
		{
			name: "case 1.2.2 negative nodes",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: -7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: -13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: -1,
						},
					},
				},
			},
			targetSum: 0,
			want:      true,
		},
		{
			name: "case 1.2.3 negative nodes",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: -7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: -13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: -1,
						},
					},
				},
			},
			targetSum: 16,
			want:      true,
		},
		{
			name: "case 1.2.4 negative nodes",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: -4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: -7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: -13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: -1,
						},
					},
				},
			},
			targetSum: 3,
			want:      false,
		},
		{
			name: "case 1.3 leetcode",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			targetSum: 1,
			want:      false,
		},
		{
			name: "case 1.4 leetcode",
			root: &TreeNode{
				Val: 1,
			},
			targetSum: 1,
			want:      true,
		},
		{
			name: "case 1.5.1 leetcode",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
								Left: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			targetSum: 6,
			want:      false,
		},
		{
			name: "case 1.6.1 leetcode",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
								Left: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			targetSum: 21,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.root, tt.targetSum)
			assert.Equal(t, tt.want, got)
		})
	}
}
