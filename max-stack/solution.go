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
	return MaxStack{
		heap: maxHeap{
			indices: make(map[heapEntry]int),
		},
	}
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

func (s *MaxStack) Len() int {
	s.compact()
	return s.heap.Len()
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

	maxHeap struct {
		es      []heapEntry
		indices map[heapEntry]int
	}
)

func (e heapEntry) Less(than heapEntry) bool {
	if e.value == than.value {
		return e.index < than.index
	}

	return e.value < than.value
}
func (e heapEntry) Equals(than heapEntry) bool { return e.value == than.value && e.index == than.index }

func (h maxHeap) Len() int { return len(h.es) }
func (h maxHeap) Swap(i, j int) {
	e1 := h.es[i]
	e2 := h.es[j]
	h.es[i], h.es[j] = h.es[j], h.es[i]
	h.indices[e1] = j
	h.indices[e2] = i
}
func (h maxHeap) Less(i, j int) bool { return !h.es[i].Less(h.es[j]) }
func (h *maxHeap) Push(x any) {
	e := x.(heapEntry)
	(*h).es = append((*h).es, e)
	h.indices[e] = len((*h).es) - 1
}

func (h *maxHeap) Pop() any {
	old := (*h).es
	n := len(old)
	x := old[n-1]
	(*h).es = old[0 : n-1]
	delete((*h).indices, x)
	return x
}
func (h *maxHeap) peek() heapEntry { return (*h).es[0] }
func (h maxHeap) indexOf(e heapEntry) int {
	i, ok := h.indices[e]
	if !ok {
		return -1
	}

	return i
}
