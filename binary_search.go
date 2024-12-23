package main

import "fmt"

func main() {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	needle := 9

	index, ok := BinarySearch(haystack, needle)

	if ok == true {
		fmt.Printf("Индекс элемента - %d\n", index)
	} else {
		fmt.Println("Элемент не найден(")
	}
}

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
