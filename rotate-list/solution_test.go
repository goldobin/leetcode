package rotate_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			name: "case 0.1",
			head: nil,
			k:    0,
			want: nil,
		},
		{
			name: "case 0.2",
			head: nil,
			k:    1,
			want: nil,
		},
		{
			name: "case 0.3",
			head: &ListNode{Val: 1},
			k:    1,
			want: &ListNode{Val: 1},
		},
		{
			name: "case 0.4",
			head: &ListNode{Val: 1},
			k:    10,
			want: &ListNode{Val: 1},
		},
		{
			name: "case 0.5",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    0,
			want: makeLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case 1.1 leetcode",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: makeLinkedList([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "case 1.2",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    5,
			want: makeLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case 1.3",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    10,
			want: makeLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case 1.4",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    11,
			want: makeLinkedList([]int{5, 1, 2, 3, 4}),
		},
		{
			name: "case 1.5",
			head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			k:    14,
			want: makeLinkedList([]int{2, 3, 4, 5, 1}),
		},
		{
			name: "case 1.6 leetcode",
			head: makeLinkedList([]int{0, 1, 2}),
			k:    4,
			want: makeLinkedList([]int{2, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.head, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func makeLinkedList(es []int) *ListNode {
	if len(es) == 0 {
		return nil
	}

	head := &ListNode{Val: es[0]}
	cur := head
	for i := 1; i < len(es); i++ {
		cur.Next = &ListNode{Val: es[i]}
		cur = cur.Next
	}

	return head
}
