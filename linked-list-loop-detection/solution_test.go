package linked_list_loop_detection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type headAndCycle struct {
	head  *node
	cycle *node
}

func Test_findCycle(t *testing.T) {
	tests := []struct {
		name         string
		headAndCycle headAndCycle
	}{
		{
			name: "case 01 null head",
			headAndCycle: headAndCycle{
				head:  nil,
				cycle: nil,
			},
		},
		{
			name: "case 02 no cycle",
			headAndCycle: headAndCycle{
				head:  &node{1, &node{2, &node{3, nil}}},
				cycle: nil,
			},
		},
		{
			name:         "case 03 cycle length 0",
			headAndCycle: makeLinkedList(0, 0),
		},
		{
			name:         "case 03 cycle length 1",
			headAndCycle: makeLinkedList(0, 1),
		},
		{
			name:         "case 03 cycle length 2",
			headAndCycle: makeLinkedList(0, 2),
		},
		{
			name:         "case 03 cycle length 3",
			headAndCycle: makeLinkedList(0, 3),
		},
		{
			name:         "case 04 cycle length 0",
			headAndCycle: makeLinkedList(1, 0),
		},
		{
			name:         "case 04 cycle length 1",
			headAndCycle: makeLinkedList(1, 1),
		},
		{
			name:         "case 04 cycle length 2",
			headAndCycle: makeLinkedList(1, 2),
		},
		{
			name:         "case 04 cycle length 3",
			headAndCycle: makeLinkedList(1, 3),
		},
		{
			name:         "case 05 cycle length 0",
			headAndCycle: makeLinkedList(2, 0),
		},
		{
			name:         "case 05 cycle length 1",
			headAndCycle: makeLinkedList(2, 1),
		},
		{
			name:         "case 05 cycle length 2",
			headAndCycle: makeLinkedList(2, 2),
		},
		{
			name:         "case 05 cycle length 3",
			headAndCycle: makeLinkedList(2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCycle(tt.headAndCycle.head)
			assert.True(t, tt.headAndCycle.cycle == got)
		})
	}
}

func makeLinkedList(length int, cycleLength int) headAndCycle {
	var (
		head    = &node{0, nil}
		current = head
	)

	for i := 0; i < length-1; i++ {
		n := &node{i + 1, nil}
		current.next = n
		current = n
	}

	if cycleLength < 0 {
		return headAndCycle{
			head:  head,
			cycle: nil,
		}
	}

	cycleNode := current
	for i := 0; i < cycleLength; i++ {
		n := &node{length + i, nil}
		current.next = n
		current = n
	}

	current.next = cycleNode
	return headAndCycle{
		head:  head,
		cycle: cycleNode,
	}
}
