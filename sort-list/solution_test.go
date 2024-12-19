package sort_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "case 1.1 leetcode",
			head: makeLinkedList([]int{4, 2, 1, 3}),
			want: makeLinkedList([]int{1, 2, 3, 4}),
		},
		{
			name: "case 1.2 leetcode",
			head: makeLinkedList([]int{-1, 5, 3, 4, 0}),
			want: makeLinkedList([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "case 1.3 leetcode",
			head: makeLinkedList(nil),
			want: makeLinkedList(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(tt.head)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_split(t *testing.T) {
	tests := []struct {
		name      string
		head      *ListNode
		wantLeft  *ListNode
		wantRight *ListNode
	}{
		{
			name:      "case 1.0",
			head:      makeLinkedList(nil),
			wantLeft:  nil,
			wantRight: nil,
		},
		{
			name:      "case 1.1",
			head:      makeLinkedList([]int{1}),
			wantLeft:  makeLinkedList([]int{1}),
			wantRight: nil,
		},
		{
			name:      "case 1.2",
			head:      makeLinkedList([]int{1, 2}),
			wantLeft:  makeLinkedList([]int{1}),
			wantRight: makeLinkedList([]int{2}),
		},
		{
			name:      "case 1.3",
			head:      makeLinkedList([]int{1, 2, 3}),
			wantLeft:  makeLinkedList([]int{1}),
			wantRight: makeLinkedList([]int{2, 3}),
		},
		{
			name:      "case 1.4",
			head:      makeLinkedList([]int{1, 2, 3, 4}),
			wantLeft:  makeLinkedList([]int{1, 2}),
			wantRight: makeLinkedList([]int{3, 4}),
		},
		{
			name:      "case 1.5",
			head:      makeLinkedList([]int{1, 2, 3, 4, 5}),
			wantLeft:  makeLinkedList([]int{1, 2}),
			wantRight: makeLinkedList([]int{3, 4, 5}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right := split(tt.head)
			assert.Equal(t, tt.wantLeft, left)
			assert.Equal(t, tt.wantRight, right)
		})
	}
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		left  *ListNode
		right *ListNode
		want  *ListNode
	}{
		{
			name:  "case 1.0",
			left:  nil,
			right: nil,
			want:  nil,
		},
		{
			name:  "case 1.1",
			left:  makeLinkedList([]int{1}),
			right: nil,
			want:  makeLinkedList([]int{1}),
		},
		{
			name:  "case 1.2",
			left:  nil,
			right: makeLinkedList([]int{1}),
			want:  makeLinkedList([]int{1}),
		},
		{
			name:  "case 1.3",
			left:  makeLinkedList([]int{1}),
			right: makeLinkedList([]int{2}),
			want:  makeLinkedList([]int{1, 2}),
		},
		{
			name:  "case 1.4",
			left:  makeLinkedList([]int{2}),
			right: makeLinkedList([]int{1}),
			want:  makeLinkedList([]int{1, 2}),
		},
		{
			name:  "case 1.5",
			left:  makeLinkedList([]int{1, 3}),
			right: makeLinkedList([]int{2, 4}),
			want:  makeLinkedList([]int{1, 2, 3, 4}),
		},
		{
			name:  "case 1.6",
			left:  makeLinkedList([]int{2, 4}),
			right: makeLinkedList([]int{1, 3}),
			want:  makeLinkedList([]int{1, 2, 3, 4}),
		},
		{
			name:  "case 1.7",
			left:  makeLinkedList([]int{1, 2, 3}),
			right: makeLinkedList([]int{4, 5, 6}),
			want:  makeLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.left, tt.right)
			assert.Equal(t, tt.want, got)
		})
	}
}

func makeLinkedList(es []int) *ListNode {
	if len(es) == 0 {
		return nil
	}
	head := &ListNode{Val: es[0]}
	node := head
	for i := 1; i < len(es); i++ {
		node.Next = &ListNode{Val: es[i]}
		node = node.Next
	}
	return head
}
