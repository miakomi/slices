package slices

import (
	"golang.org/x/exp/constraints"
)

type SliceType interface {
	string | constraints.Float | constraints.Integer
}

func Search[T SliceType](a []T, num T) (bool, []int) {
  res := []int{}
  
  for i := 0; i < len(a); i++ {
    if a[i] == num {
    res = append(res, i)
    }
  }
  
  if len(res) != 0 {
    return true, res
  }
  
  return false, res
}

func Reverse[T SliceType](a []T) {
  left := 0
  right := len(a) - 1
  
  for left < right {
    a[left], a[right] = a[right], a[left]
    
    left ++
    right --
  }
}