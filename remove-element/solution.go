package remove_element

func removeElement(nums []int, val int) int {

	var (
		slow = len(nums) - 1
		fast = len(nums) - 1
	)

	for fast >= 0 {
		if nums[fast] != val {
			fast--
			continue
		}

		if fast != slow {
			nums[fast] = nums[slow]
		}

		slow--
		fast--
	}

	return slow + 1
}
