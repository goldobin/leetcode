package path_sum

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	isLeaf := root.Left == nil && root.Right == nil
	newTargetSum := targetSum - root.Val
	if isLeaf && newTargetSum == 0 {
		return true
	}

	if hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum) {
		return true
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
