package rotate_array

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	var (
		i     = 0
		e     = nums[i]
		n     = len(nums)
		cycle = 0
	)

	for j := 0; j < n; j++ {
		i = (i + k) % n
		buf := nums[i]
		nums[i] = e
		if i == cycle {
			if j == n-1 {
				break
			}
			cycle += 1
			i = cycle
			buf = nums[i]
		}
		e = buf
	}
}
