package quad_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_construct(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want *Node
	}{
		{
			name: "case 01 leetcode",
			in: [][]int{
				{0, 1},
				{1, 0},
			},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    false,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
		{
			name: "case 02 leetcode",
			in: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    false,
					IsLeaf: false,
					TopLeft: &Node{
						Val:    false,
						IsLeaf: true,
					},
					TopRight: &Node{
						Val:    false,
						IsLeaf: true,
					},
					BottomLeft: &Node{
						Val:    true,
						IsLeaf: true,
					},
					BottomRight: &Node{
						Val:    true,
						IsLeaf: true,
					},
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
		{
			name: "case 03.1 leetcode",
			in: [][]int{
				{0},
			},
			want: &Node{
				Val:    false,
				IsLeaf: true,
			},
		},
		{
			name: "case 03.2 leetcode",
			in: [][]int{
				{1},
			},
			want: &Node{
				Val:    true,
				IsLeaf: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := construct(tt.in)
			if !assert.NotNil(t, got) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
