package main

import "fmt"

/*
Binary Search: rather than eliminating one element at the time, we eliminate half of them at the time.
only works on sorted data.

and it has BigO:

Best: O(1) if the middle is the target
Worst: O(log n) based on "n" as input it will take "log n" steps to run, example "16" elements it will take 4 steps

*/

func main() {
	items := []int{1, 2, 3, 10, 20, 24, 30, 32, 40, 50, 55, 59, 100}
	fmt.Println(binarySearch(items, 30))
}

func binarySearch(items []int, target int) int {
	left := 0
	right := len(items) - 1

	for left <= right {
		middle := (left + right) / 2
		current := items[middle]
		if current == target {
			return middle
		}

		if current < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
