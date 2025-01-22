package kth_smallest_element_in_a_bst

func kthSmallest(root *TreeNode, k int) int {
	if k < 1 {
		return -1
	}

	_, v := iterate(root, k)
	return v
}

func iterate(n *TreeNode, k int) (int, int) {
	if n == nil {
		return 0, -1
	}

	li, lv := iterate(n.Left, k)
	if li >= k {
		return li, lv
	}

	if li+1 == k {
		return k, n.Val
	}

	rightK := k - li - 1
	ri, rv := iterate(n.Right, rightK)
	return li + ri + 1, rv
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
