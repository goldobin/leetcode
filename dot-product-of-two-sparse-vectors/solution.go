package dot_product_of_two_sparse_vectors

type SparseVector struct {
	nums           []int
	nonZeroIndices map[int]struct{}
}

func Constructor(nums []int) SparseVector {
	var nonZeroIndicesCount int
	for _, n := range nums {
		if n == 0 {
			continue
		}

		nonZeroIndicesCount++
	}

	nonZeroIndices := make(map[int]struct{}, nonZeroIndicesCount)
	for i, n := range nums {
		if n == 0 {
			continue
		}

		nonZeroIndices[i] = struct{}{}
	}

	return SparseVector{
		nums:           nums,
		nonZeroIndices: nonZeroIndices,
	}
}

// Return the dotProduct of two sparse vectors
func (v *SparseVector) dotProduct(otherV SparseVector) int {
	if len(v.nums) != len(otherV.nums) {
		panic("vectors lengths do not match")
	}

	var v1 SparseVector
	var v2 SparseVector

	if len(v.nonZeroIndices) < len(otherV.nonZeroIndices) {
		v1, v2 = *v, otherV
	} else {
		v1, v2 = otherV, *v
	}

	var result int
	for k := range v1.nonZeroIndices {
		if _, ok := v2.nonZeroIndices[k]; !ok {
			continue
		}

		result += v1.nums[k] * v2.nums[k]
	}

	return result
}
