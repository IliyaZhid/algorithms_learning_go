package binary_search

// LeftBinarySearch ищет первое вхождение target в сортированном слайсе
func LeftBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	if left > target || right > target {
		return -1
	}

	for left+1 < right {
		middle := (left + right) / 2

		if nums[middle] < target {
			left = middle
		} else {
			right = middle
		}
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	return -1
}
