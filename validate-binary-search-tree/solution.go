package validate_binary_search_tree

import "math"

func isValidBST(root *TreeNode) bool {
	return iterate(root, math.MinInt, math.MaxInt)
}

func iterate(n *TreeNode, leftBound, rightBound int) bool {
	if n == nil {
		return true
	}

	if n.Val <= leftBound || n.Val >= rightBound {
		return false
	}

	return iterate(n.Left, leftBound, n.Val) && iterate(n.Right, n.Val, rightBound)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
