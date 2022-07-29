package slices

import "golang.org/x/exp/constraints"

type SliceType interface {
  string | constraints.Float | constraints.Integer
}

func Reverse[T SliceType] (a []T) {
	left := 0
	right := len(a) -1

	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}
