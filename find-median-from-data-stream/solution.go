package find_median_from_data_stream

type MedianFinder struct{}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
}

func (mf *MedianFinder) FindMedian() float64 {
	return 0
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

		swap(h.es, i, pi)
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
			swap(h.es, i, li)
			i = li
		} else {
			swap(h.es, i, ri)
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

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
