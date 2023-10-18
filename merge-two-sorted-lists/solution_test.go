package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 */

func Test_MergeTwoLists(t *testing.T) {

	cases := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			list1:    nil,
			list2:    &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			list1:    &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			list2:    nil,
			expected: &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
		},

		{
			list1:    &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			list2:    &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			expected: &ListNode{0, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}}},
		},
		{
			list1:    &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			list2:    &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			expected: &ListNode{0, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}}},
		},
	}

	for i, c := range cases {
		result := mergeTwoLists(c.list1, c.list2)

		if !reflect.DeepEqual(c.expected, result) {
			t.Errorf("case %d: result doesnt equal expected result, expected %v got %v", i+1, c.expected, result)
		}
	}

}
