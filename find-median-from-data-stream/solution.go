package find_median_from_data_stream

import (
	"container/heap"
)

type MedianFinder struct {
	left  maxHeap
	right minHeap
}

func Constructor() MedianFinder {
	lh := maxHeap{}
	heap.Init(&lh)
	rh := minHeap{}
	heap.Init(&rh)
	return MedianFinder{
		left:  lh,
		right: rh,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == 0 {
		heap.Push(&mf.left, num)
		return
	}

	l := mf.left.Peek()
	if num < l {
		if mf.left.Len()-1 == mf.right.Len() {
			heap.Pop(&mf.left)
			heap.Push(&mf.right, l)
		}

		heap.Push(&mf.left, num)
		return
	}

	if len(mf.right) == 0 {
		heap.Push(&mf.right, num)
		return
	}

	r := mf.right.Peek()
	if r < num {
		if mf.left.Len() == mf.right.Len() {
			heap.Pop(&mf.right)
			heap.Push(&mf.left, r)
		}
		heap.Push(&mf.right, num)
		return
	}

	if len(mf.left) == len(mf.right) {
		heap.Push(&mf.left, num)
	} else {
		heap.Push(&mf.right, num)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() == 0 && mf.right.Len() == 0 {
		return 0.0
	}

	l := mf.left.Peek()
	if mf.left.Len() != mf.right.Len() {
		return float64(l)
	} else {
		r := mf.right.Peek()
		return float64(l+r) / 2.0
	}
}

type (
	minHeap []int
	maxHeap []int
)

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minHeap) Peek() int {
	return (*h)[0]
}

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *maxHeap) Peek() int {
	return (*h)[0]
}
