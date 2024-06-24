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
	fmt.Println(binarySearchRecursive(items, 30))
}

// O(log n) Time | O(1) Space
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

// O(log n) Time | O(log n) Space
func binarySearchRecursive(items []int, target int) int {
	return recursive(items, target, 0, len(items)-1)
}

func recursive(items []int, target int, left int, right int) int {
	//base case
	if left > right {
		return -1
	}
	//calculate middle
	mid := (left + right) / 2
	candidate := items[mid]
	if candidate == target {
		return mid
	}
	if candidate > target {
		return recursive(items, target, left, mid-1)
	} else {
		return recursive(items, target, mid+1, right)
	}
}
