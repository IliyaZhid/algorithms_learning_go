package two_pointers

// RemoveAllOccurrences удаляет все вхождения target из массива, сохраняя порядок остальных элементов.
// Использует алгоритм двух указателей.
// Пример: arr = [3, 2, 2, 3], target = 3 → [2, 2]
func RemoveAllOccurrences(arr []int, target int) []int {
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] != target {
			arr[left] = arr[right]
			left++
		}
	}

	return arr[:left] // Возвращаем срез без удалённых элементов
}
