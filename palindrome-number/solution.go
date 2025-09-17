package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var (
		nums = make([]byte, 0, 10) // max 10 digits in 2^31 integer
		d1   = 1
	)
	for i := 0; i < 10; i++ {
		if x/d1 == 0 {
			break
		}

		var (
			d2  = d1 * 10
			rem = x % d2
			v   = rem / d1
		)
		nums = append(nums, byte(v))
		d1 = d2
	}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}

	return true
}
