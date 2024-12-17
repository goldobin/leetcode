package linked_list_cycle

func hasCycle(head *ListNode) bool {
	p1, p2 := head, head

	for {
		if p1 != nil {
			p1 = p1.Next
		} else {
			return false
		}

		if p2 != nil && p2.Next != nil {
			p2 = p2.Next.Next
		} else {
			return false
		}

		if p1 == p2 {
			return true
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
