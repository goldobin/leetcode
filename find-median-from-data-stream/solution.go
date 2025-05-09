package find_median_from_data_stream

type MedianFinder struct {
	left  heap
	right heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  newMaxHeap(),
		right: newMinHeap(),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	l, lOK := mf.left.peek()
	r, rOK := mf.right.peek()

	if !lOK {
		mf.left.push(num)
		return
	}

	if num < l {
		if mf.left.size()-1 == mf.right.size() {
			_, _ = mf.left.pull()
			mf.right.push(l)
		}

		mf.left.push(num)
		return
	}

	if !rOK {
		mf.right.push(num)
		return
	}

	if r < num {
		if mf.left.size() == mf.right.size() {
			_, _ = mf.right.pull()
			mf.left.push(r)
		}
		mf.right.push(num)
		return
	}

	if mf.left.size() == mf.right.size() {
		mf.left.push(num)
	} else {
		mf.right.push(num)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.size() == 0 && mf.right.size() == 0 {
		return 0.0
	}

	l, _ := mf.left.peek()
	if mf.left.size() != mf.right.size() {
		return float64(l)
	} else {
		r, _ := mf.right.peek()
		return float64(l+r) / 2.0
	}
}

type (
	heap struct {
		es      []int
		compare compareFn
	}
	compareFn func(int, int) int
)

func greater(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func less(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return -1
}

func newMinHeap() heap {
	return heap{
		compare: less,
	}
}

func newMaxHeap() heap {
	return heap{
		compare: greater,
	}
}

func (h *heap) push(v int) {
	h.es = append(h.es, v)
	i := len(h.es) - 1

	for i > 0 {
		pi := i / 2
		if h.compare(h.es[i], h.es[pi]) < 0 {
			break
		}

		h.es[i], h.es[pi] = h.es[pi], h.es[i]
		i = pi
	}
}

func (h *heap) pull() (int, bool) {
	if len(h.es) == 0 {
		return 0, false
	}

	r := h.es[0]
	h.es[0] = h.es[len(h.es)-1]
	h.es = h.es[:len(h.es)-1]

	i := 0

	for i < len(h.es)/2 {
		li := i * 2
		ri := li + 1

		if !(h.compare(h.es[i], h.es[li]) < 0 || h.compare(h.es[i], h.es[ri]) < 0) {
			break
		}

		if h.compare(h.es[li], h.es[ri]) > 0 {
			h.es[i], h.es[li] = h.es[li], h.es[i]
			i = li
		} else {
			h.es[i], h.es[ri] = h.es[ri], h.es[i]
			i = ri
		}
	}

	return r, true
}

func (h *heap) peek() (int, bool) {
	if len(h.es) == 0 {
		return 0, false
	}
	return (h.es)[0], true
}

func (h *heap) size() int {
	return len(h.es)
}
