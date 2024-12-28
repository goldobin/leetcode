package rotate_list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if k == 0 {
		return head
	}

	_, _, newHead := iterate(head, k, head, 1)
	return newHead
}

func iterate(head *ListNode, k int, node *ListNode, l int) (int, *ListNode, *ListNode) {
	if node.Next == nil {
		b := k % l
		if b == 0 {
			return 0, nil, head
		}

		return b, node, nil
	}

	n, last, result := iterate(head, k, node.Next, l+1)

	if n == 0 {
		return 0, nil, result
	}

	n -= 1

	if n == 0 {
		newHead := node.Next
		node.Next = nil
		last.Next = head
		return 0, nil, newHead
	}

	return n, last, nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
