package climb_stairs_recursive

func climbStairs(n int) int {
	memory := make([]int, n+1)
	memory[0] = 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			memory[i] = 1
		} else if i == 2 {
			memory[i] = 2
		} else {
			memory[i] = memory[i-1] + memory[i-2]
		}
	}

	return memory[n]
}
