package linear_search

// LastEven Возвращает последний четный элемент слайса
func LastEven(nums []int) int {
	lastEven := -1

	for _, val := range nums {
		if val%2 == 0 {
			lastEven = val
		}
	}

	return lastEven
}
