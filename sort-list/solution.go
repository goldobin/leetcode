package sort_list

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)

	left = sortList(left)
	right = sortList(right)

	head = merge(left, right)

	return head
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	if head1.Val < head2.Val {
		head1.Next = merge(head1.Next, head2)
		return head1
	}

	head2.Next = merge(head1, head2.Next)
	return head2
}

func split(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	return head, slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
