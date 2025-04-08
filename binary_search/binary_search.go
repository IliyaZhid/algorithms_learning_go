package binary_search

func BinarySearch(nums []int, target int) (int, bool) {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid, true
		}

		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, false
}
