package binary_search

// CopyTime находит минимальное время копирования на двух принтерах с разной скоростью печати
func CopyTime(n, x, y int) int {
	left := 0
	right := (n - 1) * max(x, y)

	for left-1 < right {
		middle := (left + right) / 2
		if middle/x+middle/y < n-1 {
			left = middle
		} else {
			right = middle
		}
	}

	return right + min(x, y)
}
