package remove_duplicates_from_sorted_list_ii

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		firstNodeParent *ListNode = nil
		firstNode                 = head
		lastNode                  = head
	)

	for {
		lastNodeChild := lastNode.Next
		if lastNodeChild != nil && lastNode.Val == lastNodeChild.Val {
			lastNode = lastNodeChild
			continue
		}

		if firstNode == lastNode {
			firstNodeParent = firstNode
		} else if firstNodeParent != nil {
			firstNodeParent.Next = lastNodeChild
		} else {
			head = lastNodeChild
		}

		if lastNodeChild == nil {
			break
		}

		firstNode = lastNodeChild
		lastNode = lastNodeChild
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
