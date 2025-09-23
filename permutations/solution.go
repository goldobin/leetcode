package permutations

var factorials = []int{
	1,
	1,
	2,
	2 * 3,
	2 * 3 * 4,
	2 * 3 * 4 * 5,
	2 * 3 * 4 * 5 * 6,
}

func permute(nums []int) [][]int {
	if len(nums) > len(factorials)-1 {
		panic("too many numbers")
	}

	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	var (
		mask      = make([]bool, len(nums))
		resultLen = factorials[len(nums)]
		result    = make([][]int, resultLen)
	)
	for i := 0; i < resultLen; i++ {
		result[i] = make([]int, 0, len(nums))
	}

	iterate(result, nums, mask, len(nums))
	return result
}

func iterate(dst [][]int, nums []int, mask []bool, l int) {
	if l == 1 {
		for i, n := range nums {
			if mask[i] {
				continue
			}
			dst[0] = append(dst[0], n)
		}
		return
	}

	var (
		shardLen = factorials[l-1]
		shardIx  = 0
	)
	for i, n := range nums {
		if mask[i] {
			continue
		}

		mask[i] = true
		shard := dst[shardIx*shardLen : (shardIx+1)*shardLen]
		for j := 0; j < len(shard); j++ {
			shard[j] = append(shard[j], n)
		}
		iterate(shard, nums, mask, l-1)
		mask[i] = false
		shardIx += 1
	}
}
