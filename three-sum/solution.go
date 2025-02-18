package three_sum

import (
	"bytes"
	"cmp"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"slices"
	"sort"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0, len(nums))
	i := 0
	for i < len(nums)-2 {
		pivot := nums[i]
		searchSpace := nums[i+1:]
		result = append2Sums(result, searchSpace, pivot)
		for {
			i += 1
			if i >= len(nums) || pivot != nums[i] {
				break
			}
		}
	}

	return result
}

func append2Sums(dst [][]int, nums []int, pivot int) [][]int {
	i := 0
	for i < len(nums)-1 {
		var (
			vi          = nums[i]
			searchSpace = nums[i+1:]
			vj          = -pivot - vi
		)

		_, found := sort.Find(len(searchSpace), func(j int) int {
			return cmp.Compare(vj, searchSpace[j])
		})

		if found {
			var triad []int
			if pivot < 0 {
				triad = []int{pivot, vi, vj}
			} else {
				triad = []int{vi, vj, pivot}
			}

			dst = append(dst, triad)
		}

		for {
			i += 1
			if i >= len(nums) || vi != nums[i] {
				break
			}
		}
	}
	return dst
}

func threeSumBruteForce(nums []int) [][]int {
	result := make([][]int, 0)
	hexes := make(map[string]struct{})
	hexer := newTripletHexer()

	var t triplet
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				t[0] = nums[i]
				t[1] = nums[j]
				t[2] = nums[k]
				sum := 0

				for _, v := range t {
					sum += v
				}

				if sum != 0 {
					continue
				}

				slices.Sort(t[:])
				s := hexer.EncodeToString(t)
				_, ok := hexes[s]
				if ok {
					continue
				}

				hexes[s] = struct{}{}
				result = append(result, slices.Clone(t[:]))
			}
		}
	}

	return result
}

type (
	triplet      [3]int
	tripletHexer struct {
		buff bytes.Buffer
	}
)

func newTripletHexer() tripletHexer {
	var v tripletHexer
	v.buff.Grow(4 * 3)
	return v
}

func (h *tripletHexer) EncodeToString(t triplet) string {
	buff := &h.buff
	buff.Reset()
	for _, n := range t {
		if err := binary.Write(buff, binary.NativeEndian, int32(n)); err != nil {
			panic(fmt.Sprintf("unexpected error %v", err))
		}
	}

	return hex.EncodeToString(buff.Bytes())
}
