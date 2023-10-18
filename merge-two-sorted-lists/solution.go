package merge_two_sorted_lists

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head1 := list1
	head2 := list2

	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	var (
		result         *ListNode = nil
		resultLastNode *ListNode = nil
	)

	if head1.Val < head2.Val {
		result = head1
		resultLastNode = head1
		head1 = head1.Next
	} else {
		result = head2
		resultLastNode = head2
		head2 = head2.Next
	}

	for {
		if head1 == nil {
			resultLastNode.Next = head2
			break
		}

		if head2 == nil {
			resultLastNode.Next = head1
			break
		}

		if head1.Val < head2.Val {
			resultLastNode.Next = head1
			resultLastNode = head1
			head1 = head1.Next

		} else {
			resultLastNode.Next = head2
			resultLastNode = head2
			head2 = head2.Next
		}
	}

	return result
}
