package slices

import (
	"golang.org/x/exp/constraints"
)

type SliceType interface {
	string | constraints.Float | constraints.Integer
}

type Numbers interface {
	constraints.Float | constraints.Integer
}

func Search[T Numbers](a []T, num T) (bool, int) {
	left := 0
	right := len(a) - 1
	mid := 0

	for left <= right {
		mid = ((right - left) / 2) + left

		if num == a[mid] {
			return true, mid
		} else if num < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false, -1
}
