package two_pointers

// MinAbsDiff Находит пару соседних элементов с минимальной абсолютной разницей
func MinAbsDiff(nums []int) (int, int) {
	if len(nums) < 2 {
		return -1, -1
	}

	minA, minB := 0, 1

	for i := 0; i < len(nums)-1; i++ {
		diff := getAbsDiff(nums[i], nums[i+1])

		if diff < getAbsDiff(nums[minA], nums[minB]) {
			minA, minB = i, i+1
		}
	}

	return minA, minB
}

func getAbsDiff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
