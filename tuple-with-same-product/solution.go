package tuple_with_same_product

import (
	"fmt"
	"slices"
)

func tupleSameProduct(nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	all := make(map[int]struct{}, len(nums))
	counted := make(map[key]struct{}, len(nums))

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		all[nums[i]] = struct{}{}
	}

	return count(counted, all, nums)
}

func count(counted map[key]struct{}, all map[int]struct{}, nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	n1 := nums[0]
	n2 := nums[len(nums)-1]
	k := key{n1, n2}
	prod := n1 * n2

	if _, ok := counted[k]; ok {
		return 0
	}

	counted[k] = struct{}{}

	n := 0
	for _, n3 := range nums[1 : len(nums)-1] {
		if prod%n3 != 0 {
			continue
		}

		n4 := prod / n3
		if n4 <= n3 || n4 >= n2 {
			continue
		}

		if _, ok := all[n4]; !ok {
			continue
		}
		n += 8
		fmt.Printf("{%d, %d, %d, %d}\n", n1, n2, n3, n4)
	}

	n += count(counted, all, nums[0:len(nums)-1])
	n += count(counted, all, nums[1:])

	return n
}

type key struct {
	a, b int
}
