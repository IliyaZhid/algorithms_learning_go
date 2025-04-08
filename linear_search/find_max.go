package linear_search

// FindMax находит максимальное число в массиве.
// Пример: [5, 1, 44, 100, 12] → 100
func FindMax(arr []int) int {

	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}

	return m
}
