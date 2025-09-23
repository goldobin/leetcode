package symmetric_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const NoVal = -101

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "case 0.1 nil node",
			want: true,
		},
		{
			name: "case 0.2 node without children",
			root: buildTree(1),
			want: true,
		},
		{
			name: "case 0.3.1 node with one child",
			root: buildTree(1, 2),
			want: false,
		},
		{
			name: "case 0.3.2 node with one child",
			root: buildTree(1, NoVal, 3),
			want: false,
		},
		{
			name: "case 1.1 node with 2 unequal children",
			root: buildTree(1, 2, 3),
			want: false,
		},
		{
			name: "case 1.2 node with 2 equal children",
			root: buildTree(1, 2, 2),
			want: true,
		},
		{
			name: "case 1.3.1 non symmetric tree",
			root: buildTree(1, 2, 2, 4, 3, 3, 5),
			want: false,
		},
		{
			name: "case 1.3.2 non symmetric tree",
			root: buildTree(1, 2, 2, 5, NoVal, 3, 5),
			want: false,
		},
		{
			name: "case 1.3.3 non symmetric tree",
			root: buildTree(1, 2, 2, 4, 3, NoVal, 4),
			want: false,
		},
		{
			name: "case 1.3.4 symmetric tree",
			root: buildTree(1, 2, 2, 4, 3, 3, 4),
			want: true,
		},
		{
			name: "case 2.1 leetcode",
			root: buildTree(1, 2, 2, 3, 4, 4, 3),
			want: true,
		},
		{
			name: "case 2.2 leetcode",
			root: buildTree(1, 2, 2, NoVal, 3, NoVal, 3),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}

func buildTree(vs ...int) *TreeNode {
	return buildTreeNode(0, vs)
}

func buildTreeNode(i int, vs []int) *TreeNode {
	if i >= len(vs) {
		return nil
	}

	v := vs[i]
	if v == NoVal {
		return nil
	}
	li := i*2 + 1
	ri := i*2 + 2

	return &TreeNode{
		Val:   v,
		Left:  buildTreeNode(li, vs),
		Right: buildTreeNode(ri, vs),
	}
}
