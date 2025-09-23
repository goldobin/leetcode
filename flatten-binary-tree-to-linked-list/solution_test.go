package flatten_binary_tree_to_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const NoVal = -101

func Test_flatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "case 0.1 nil root",
		},
		{
			name: "case 0.2 single root",
			root: arrayToTree(1),
			want: arrayToTree(1),
		},
		{
			name: "case 0.3 only left node",
			root: arrayToTree(1, 2),
			want: arrayToTree(1, NoVal, 2),
		},
		{
			name: "case 0.4 only right node",
			root: arrayToTree(1, NoVal, 3),
			want: arrayToTree(1, NoVal, 3),
		},
		{
			name: "case 0.5 two left nodes",
			root: arrayToTree(1, 2, NoVal, 3, NoVal),
			want: arrayToTree(1, NoVal, 2, NoVal, 3),
		},
		{
			name: "case 0.6 two right nodes",
			root: arrayToTree(1, NoVal, 2, NoVal, 3),
			want: arrayToTree(1, NoVal, 2, NoVal, 3),
		},
		{
			name: "case 1.1 balanced tree",
			root: arrayToTree(1, 2, 5, 3, 4, 6, 7),
			want: arrayToTree(
				1,
				NoVal, 2,
				NoVal, 3,
				NoVal, 4,
				NoVal, 5,
				NoVal, 6,
				NoVal, 7,
			),
		},
		{
			name: "case 2.1 leetcode",
			root: arrayToTree(1, 2, 5, 3, 4, NoVal, 6),
			want: arrayToTree(1, NoVal, 2, NoVal, 3, NoVal, 4, NoVal, 5, NoVal, 6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cloneTree(tt.root)
			flatten(got)
			assert.Equal(t, treeToArray(tt.want), treeToArray(got))
		})
	}
}

func cloneTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	return &TreeNode{
		Val:   node.Val,
		Left:  cloneTree(node.Left),
		Right: cloneTree(node.Right),
	}
}

func arrayToTree(vs ...int) *TreeNode {
	if len(vs) == 0 || vs[0] == NoVal {
		return nil
	}

	root := &TreeNode{Val: vs[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vs) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vs) && vs[i] != NoVal {
			node.Left = arrayToTree(vs[i])
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vs) && vs[i] != NoVal {
			node.Right = arrayToTree(vs[i])
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func treeToArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, NoVal)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	return result
}
