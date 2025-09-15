package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	l, r, pivot := binarySearch(nums, target)
	result := make([]int, 2)

	if pivot == 0 || pivot > 0 && nums[pivot-1] != target {
		result[0] = pivot
	} else {
		result[0] = binarySearchL(nums, l, pivot, target)
	}

	if pivot == len(nums)-1 || pivot < len(nums)-1 && nums[pivot+1] != target {
		result[1] = pivot
	} else {
		result[1] = binarySearchR(nums, pivot+1, r, target)
	}

	return result
}

func binarySearch(nums []int, target int) (l, r, pivot int) {
	l, r = 0, len(nums)
	for l < r {
		pivot = l + (r-l)/2
		v := nums[pivot]
		if v == target {
			return l, r, pivot
		}

		if v < target {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	pivot = -1
	return l, r, pivot
}

func binarySearchL(nums []int, l, r, target int) int {
	for l < r {
		pivot := l + (r-l)/2
		v := nums[pivot]
		if v == target {
			if pivot == 0 || pivot > 0 && nums[pivot-1] != target {
				return pivot
			}

			r = pivot
			continue
		}

		if v < target {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	return -1
}

func binarySearchR(nums []int, l, r, target int) int {
	for l < r {
		pivot := l + (r-l)/2
		v := nums[pivot]
		if v == target {
			if pivot == len(nums)-1 || pivot < len(nums)-1 && nums[pivot+1] != target {
				return pivot
			}

			l = pivot + 1
			continue
		}

		if v < target {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	return -1
}
