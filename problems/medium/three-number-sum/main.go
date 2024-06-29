package main

import (
	"fmt"
	"slices"
)

//it is just math equation: Result = CurrentNumber + X + Y
//so "result" is given to us, and current number is just the first loop and x and y are 2 variables in this equation
//so for each "current number" we need to look for 2 numbers in the rest of the list

func main() {
	numbers := []int{12, 3, 1, 2, -6, 5, -8, 6}
	fmt.Println(threeNumberSum(numbers, 0))
}

func threeNumberSum(numbers []int, target int) [][]int {
	var results [][]int

	//sort it first
	slices.Sort(numbers)

	//we are looking for 3 numbers so the last current must be (length -2 )
	for idx := 0; idx < len(numbers)-2; idx++ {
		current := numbers[idx]
		// target = current + left + right
		left := idx + 1
		right := len(numbers) - 1

		for left < right {
			result := current + numbers[left] + numbers[right]
			if result == target {
				results = append(results, []int{current, numbers[left], numbers[right]})
				//move both pointer
				left++
				right--
			} else if result < target {
				//move the left forward
				left++
			} else {
				right--
			}
		}
	}
	return results
}
