package binary_search

// SearchSqrt Находит корень target или ближайшее целое
func SearchSqrt(target int) int {

	left, right := 1, target

	for left <= right {
		middle := (left + right) / 2
		square := middle * middle

		if square < target {
			left = middle + 1
			continue
		}
		if square > target {
			right = middle - 1
			continue
		}

		return middle
	}

	return right
}
