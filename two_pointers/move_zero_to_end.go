package two_pointers

// MoveZerosToEnd перемещает все нули в конец массива, сохраняя порядок остальных элементов.
// Использует алгоритм двух указателей.
// Пример: [0, 1, 0, 3, 12] → [1, 3, 12, 0, 0]
func MoveZerosToEnd(arr []int) []int {
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] != 0 {
			arr[left] = arr[right]
			left++
		}
	}

	for left < len(arr) {
		arr[left] = 0
		left++
	}

	return arr
}
