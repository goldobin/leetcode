package longest_increasing_subsequence_dynamic

import (
	"cmp"
	"slices"
)

type (
	indexEntry struct {
		val   int
		index int
	}
	index []indexEntry
)

func (ix index) fill(nums []int) {
	for i := 0; i < len(ix); i++ {
		ix[i].index = i
		ix[i].val = nums[i]
	}
	ix.sort()
}

func (ix index) sort() {
	slices.SortFunc(ix, func(a, b indexEntry) int { return cmp.Compare(a.val, b.val) })
}

func (ix index) find(v int) (int, bool) {
	return slices.BinarySearchFunc(ix, v, func(e indexEntry, t int) int {
		return cmp.Compare(e.val, t)
	})
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		memory = make([]int, len(nums))
		ix     = make(index, len(nums))
	)
	ix.fill(nums)

	result := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		sortedIx, found := ix.find(v)
		if !found {
			panic("not found")
		}

		maxSeqLen := 0
		for j := 0; j < sortedIx; j++ {
			e := ix[j]
			seqLen := memory[e.index] + 1
			if maxSeqLen < seqLen {
				maxSeqLen = seqLen
			}
		}
		if maxSeqLen < 1 {
			maxSeqLen = 1
		}
		memory[i] = maxSeqLen
		if result < maxSeqLen {
			result = maxSeqLen
		}
	}

	return result
}
