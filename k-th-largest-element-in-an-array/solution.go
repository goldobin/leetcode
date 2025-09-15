package k_th_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	if k == 0 {
		panic("k should be at least 1")
	}

	if len(nums) < k {
		panic("k out of range")
	}

	h := newMaxHeap(len(nums))
	for _, v := range nums {
		h.push(v)
	}

	var result int
	for i := 0; i < k; i++ {
		result, _ = h.pull()
	}

	return result
}

type heap struct {
	es []int
}

func newMaxHeap(capacity int) heap {
	return heap{es: make([]int, 0, capacity)}
}

func (h *heap) push(v int) {
	h.es = append(h.es, v)
	i := len(h.es) - 1

	for i > 0 {
		pi := i / 2
		if h.es[i] < h.es[pi] {
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

		if h.es[i] >= h.es[li] && h.es[i] >= h.es[ri] {
			break
		}

		if h.es[li] > h.es[ri] {
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
