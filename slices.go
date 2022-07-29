package slices

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type SliceType interface {
	string | constraints.Float | constraints.Integer
}

func SearchInt(a []int, num int) string {
	left := 0
	right := len(a) - 1
	mid := 0

	for left <= right {
		mid = ((right - left) / 2) + left

		if num == a[mid] {
			return fmt.Sprint("Ok. i = ", mid)
		} else if num < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return "NO"
}

func SearchFloat32(a []float32, num float32) string {
	left := 0
	right := len(a) - 1
	mid := 0

	for left <= right {
		mid = ((right - left) / 2) + left

		if num == a[mid] {
			return fmt.Sprint("Ok. i = ", mid)
		} else if num < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return "NO"
}

func Reverse[T SliceType](a []T) {
	left := 0
	right := len(a) - 1

	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}
