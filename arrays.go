package arrays

func Reverse(a []int) []int {
	left := 0
	right := len(a) - 1
	temp := 0

	for left < right {
		temp = a[left]
		a[left] = a[right]
		a[right] = temp

		left++
		right--
	}
	return a
}
