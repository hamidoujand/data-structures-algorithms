package main

import "fmt"

/*

Linear Search: used when data is not sorted, and its Big O:

Best Case: O(1) if data we looking for is at index 0
Worst Case: O(n)


*/

func main() {
	items := []int{1000, 2, 102, 132, 300, 30, -1, 0, 12}
	fmt.Println(linearSearch(items, -1))
}

func linearSearch(items []int, target int) int {
	for i := 0; i < len(items); i++ {
		if items[i] == target {
			return i
		}
	}
	return -1
}
