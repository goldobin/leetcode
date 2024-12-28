package invert_tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l, r := root.Left, root.Right

	root.Right = invertTree(l)
	root.Left = invertTree(r)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
