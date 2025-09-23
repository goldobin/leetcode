package flatten_binary_tree_to_linked_list

func flatten(root *TreeNode) {
	iterate(root)
}

func iterate(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	l, r := node.Left, node.Right
	if l == nil && r == nil {
		return node
	}

	if r == nil {
		lTail := iterate(node.Left)
		node.Left = nil
		node.Right = l
		return lTail
	}

	if l == nil {
		rTail := iterate(node.Right)
		return rTail
	}

	lTail, rTail := iterate(node.Left), iterate(node.Right)
	node.Left = nil
	node.Right = l
	lTail.Right = r

	return rTail
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
