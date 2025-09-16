package remove_duplicates_from_sorted_list_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{name: "case 0.1 empty", head: nil, want: nil},
		{name: "case 0.2 one node", head: &ListNode{Val: 1}, want: &ListNode{Val: 1}},
		{
			name: "case 0.3 two distinct nodes",
			head: toLinkedList(1, 2),
			want: toLinkedList(1, 2),
		},
		{
			name: "case 0.4 two nodes with same value",
			head: toLinkedList(2, 2),
			want: nil,
		},
		{
			name: "case 0.5 three nodes with two nodes at the beginning have the same value",
			head: toLinkedList(2, 2, 3),
			want: toLinkedList(3),
		},
		{
			name: "case 0.6 three nodes with two nodes at the end have the same value",
			head: toLinkedList(1, 2, 2),
			want: toLinkedList(1),
		},
		{
			name: "case 0.7 three nodes have the same value",
			head: toLinkedList(3, 3, 3),
			want: nil,
		},
		{
			name: "case 0.8 four nodes have the same value",
			head: toLinkedList(4, 4, 4, 4),
			want: nil,
		},
		{
			name: "case 0.9 four nodes with two nodes in the middle have the same value",
			head: toLinkedList(1, 2, 2, 3),
			want: toLinkedList(1, 3),
		},
		{
			name: "case 1.1 leetcode",
			head: toLinkedList(1, 2, 3, 3, 4, 4, 5),
			want: toLinkedList(1, 2, 5),
		},
		{
			name: "case 1.3 leetcode",
			head: toLinkedList(1, 1, 1, 2, 3),
			want: toLinkedList(2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			assert.Equal(t, tt.want, got)
		})
	}
}

func toLinkedList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var (
		head = &ListNode{Val: vals[0]}
		cur  = head
	)

	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}

	return head
}
