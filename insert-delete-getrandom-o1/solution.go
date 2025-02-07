package insert_delete_getrandom_o1

import (
	"math/rand/v2"
)

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		a: make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	_, ok := s.m[val]
	if ok {
		return false
	}

	newI := len(s.a)
	s.a = append(s.a, val)
	s.m[val] = newI
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	i, ok := s.m[val]
	if !ok {
		return false
	}

	delete(s.m, val)
	if len(s.a) > 1 && i != len(s.a)-1 {
		w := s.a[len(s.a)-1]
		s.a[i] = w
		s.m[w] = i
	}
	s.a = s.a[:len(s.a)-1]

	return true
}

func (s *RandomizedSet) GetRandom() int {
	i := rand.IntN(len(s.a))
	return s.a[i]
}
