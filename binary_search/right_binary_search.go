package binary_search

// RightBinarySearch ищет последнее вхождение target в сортированном слайсе
func RightBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left+1 < right {
		middle := (left + right) / 2

		if nums[middle] <= target {
			left = middle
		} else {
			right = middle
		}
	}

	if nums[right] == target {
		return right
	}

	if nums[left] == target {
		return left
	}

	return -1
}
