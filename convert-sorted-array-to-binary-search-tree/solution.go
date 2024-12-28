package convert_sorted_array_to_binary_search_tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var (
		center = len(nums) / 2
		left   = sortedArrayToBST(nums[:center])
		right  = sortedArrayToBST(nums[center+1:])
	)

	return &TreeNode{
		Val:   nums[center],
		Left:  left,
		Right: right,
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
