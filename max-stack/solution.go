package max_stack

import (
	"container/heap"
	"fmt"
)

type MaxStack struct {
	stack []stackEntry
	heap  maxHeap
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (s *MaxStack) Push(x int) {
	s.stack = append(s.stack, stackEntry{value: x})
	var (
		i  = len(s.stack) - 1
		he = heapEntry{
			value: x,
			index: i,
		}
	)

	heap.Push(&s.heap, he)
}

func (s *MaxStack) Pop() int {
	s.compact()
	if len(s.stack) == 0 {
		panic("stack is empty")
	}

	var (
		i     = len(s.stack) - 1
		value = s.stack[i].value
		he    = heapEntry{
			value: value,
			index: i,
		}
	)
	s.stack[i].thumbStone = true
	hi := s.heap.indexOf(he)

	if hi == -1 {
		panic(fmt.Sprintf("heap doesn't contain heapEntry: %v", he))
	}

	heap.Remove(&s.heap, hi)
	return value
}

func (s *MaxStack) Top() int {
	s.compact()
	if len(s.stack) == 0 {
		panic("stack is empty")
	}

	ix := len(s.stack) - 1
	return s.stack[ix].value
}

func (s *MaxStack) PeekMax() int {
	if len(s.stack) == 0 {
		panic("stack is empty")
	}

	return s.heap.peek().value
}

func (s *MaxStack) PopMax() int {
	if len(s.stack) == 0 {
		panic("stack is empty")
	}

	e := heap.Pop(&s.heap).(heapEntry)
	s.stack[e.index].thumbStone = true
	return e.value
}

func (s *MaxStack) compact() {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if !s.stack[i].thumbStone {
			break
		}
		s.stack = s.stack[:i]
	}
}

type (
	stackEntry struct {
		value      int
		thumbStone bool
	}
	heapEntry struct {
		value int
		index int
	}
	maxHeap []heapEntry
)

func (e heapEntry) Less(than heapEntry) bool {
	if e.value == than.value {
		return e.index < than.index
	}

	return e.value < than.value
}
func (e heapEntry) Equals(than heapEntry) bool { return e.value == than.value && e.index == than.index }

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return !h[i].Less(h[j]) }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(heapEntry)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *maxHeap) peek() heapEntry { return (*h)[0] }
func (h maxHeap) indexOf(e heapEntry) int {
	for i, v := range h {
		if e.Equals(v) {
			return i
		}
	}

	return -1
}
