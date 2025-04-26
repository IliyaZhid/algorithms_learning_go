package binary_search

// Diploma находит минимальную сторону квадрата в который поместятся фигуры заданных параметров
func Diploma(height, width, n int) int {
	left := max(height, width)
	right := left * n

	for left+1 < right {
		middle := (left + right) / 2
		res := (middle / width) * (middle / height)
		if res < n {
			left = middle
		} else {
			right = middle
		}
	}

	return right
}
