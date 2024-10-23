package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	counter := 0

	sums[0] = 1
	sum := 0
	for _, v := range nums {
		sum += v
		left := sum - k

		if s, ok := sums[left]; ok {
			counter += s
		}

		if s, ok := sums[sum]; ok {
			sums[sum] = s + 1
		} else {
			sums[sum] = 1
		}
	}

	return counter
}
