package rotate_image

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	t.Run("matrix 1x1", func(t *testing.T) {
		m := [][]int{
			{1},
		}

		mExpected := [][]int{
			{1},
		}
		rotate(m)

		if !reflect.DeepEqual(m, mExpected) {
			t.Errorf("expected=%v, got=%v", mExpected, m)
		}
	})

	t.Run("matrix 2x2", func(t *testing.T) {
		m := [][]int{
			{1, 2},
			{3, 4},
		}
		mExpected := [][]int{
			{3, 1},
			{4, 2},
		}
		rotate(m)

		if !reflect.DeepEqual(m, mExpected) {
			t.Errorf("expected=%v, got=%v", mExpected, m)
		}
	})

	t.Run("matrix 3x3", func(t *testing.T) {
		m := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		mExpected := [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}

		rotate(m)

		if !reflect.DeepEqual(m, mExpected) {
			t.Errorf("expected=%v, got=%v", mExpected, m)
		}
	})
	t.Run("matrix 4x4", func(t *testing.T) {
		m := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		mExpected := [][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		}

		rotate(m)

		if !reflect.DeepEqual(m, mExpected) {
			t.Errorf("expected=%v, got=%v", mExpected, m)
		}
	})
}
