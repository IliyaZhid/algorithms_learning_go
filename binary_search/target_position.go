package binary_search

// GetTargetPosition возвращает позицию элемента target или позицию на которой он должен находиться если его нет в слайсе
func GetTargetPosition(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}

	return left
}
