package majority_element

func majorityElement(nums []int) int {
	var (
		majority int
		count    int
	)

	for _, num := range nums {
		if count == 0 {
			majority = num
		}

		if num == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}
