package random_pick_with_weight

import (
	"cmp"
	"math/rand/v2"
)

type Solution struct {
	tree *node[point]
}

func Constructor(w []int) Solution {
	if len(w) == 0 {
		panic("empty array")
	}

	ps := toPoints(w)
	tree := buildTree(ps)
	return Solution{
		tree: tree,
	}
}

func (s *Solution) PickIndex() int {
	x := rand.Float64()
	n := s.tree.ceilFunc(point{x: x}, cmpPoint)

	if n == nil {
		panic("unreachable")
	}

	return n.value.index
}

type point struct {
	x     float64
	index int
}

func cmpPoint(a, b point) int {
	return cmp.Compare(a.x, b.x)
}

func toPoints(w []int) []point {
	var (
		sum    = 0
		x      = 0.0
		result = make([]point, len(w))
	)

	for _, v := range w {
		sum += v
	}

	for i, v := range w {
		result[i] = point{
			index: i,
			x:     x,
		}
		x += float64(v) / float64(sum)
	}

	return result
}

type node[E any] struct {
	value E
	left  *node[E]
	right *node[E]
}

func buildTree[E any](sortedEs []E) *node[E] {
	if len(sortedEs) == 0 {
		return nil
	}

	var (
		center = len(sortedEs) / 2
		left   = buildTree(sortedEs[:center])
		right  = buildTree(sortedEs[center+1:])
	)

	return &node[E]{
		value: sortedEs[center],
		left:  left,
		right: right,
	}
}

func (n *node[E]) ceilFunc(x E, cmp func(a, b E) int) *node[E] {
	c := cmp(x, n.value)

	if c < 0 {
		if n.left == nil {
			return nil
		}
		return n.left.ceilFunc(x, cmp)
	}

	if c > 0 {
		if n.right == nil {
			return n
		}
		result := n.right.ceilFunc(x, cmp)
		if result == nil {
			return n
		}

		return result
	}

	return n
}
