package symmetric_tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return areBranchesSymmetric(root.Left, root.Right)
}

func areBranchesSymmetric(b1 *TreeNode, b2 *TreeNode) bool {
	if b1 == nil && b2 == nil {
		return true
	}

	if b1 != nil && b2 != nil && b1.Val == b2.Val {
		return areBranchesSymmetric(b1.Left, b2.Right) && areBranchesSymmetric(b1.Right, b2.Left)
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
