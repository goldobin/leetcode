package rand10_using_rand7

import "math/rand"

func rand10() int {
	n := 0
	for i := 0; i < 10; i++ {
		n += rand7()
	}
	return n%10 + 1
}

func rand7() int {
	return rand.Intn(7) + 1
}
