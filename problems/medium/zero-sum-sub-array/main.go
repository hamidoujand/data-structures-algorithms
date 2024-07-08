package main

import "fmt"

func main() {
	numbers := []int{4, -3, 2, 4, -1, -5, 7}
	fmt.Println(zeroSumSubArray(numbers))
}

// O(n) Time | O(n) Space
func zeroSumSubArray(numbers []int) bool {
	sums := make(map[int]struct{}, len(numbers))

	//add zero for the edge case when all numbers are 0 inside of numbers
	sums[0] = struct{}{}

	currentSum := 0

	for _, num := range numbers {
		currentSum += num
		//then we check if we already had this sum inside of the set, that means
		//we have a situation like this: "sum + 0 = sum" so we have a "0" sum
		if _, ok := sums[currentSum]; ok {
			return true
		} else {
			sums[currentSum] = struct{}{}
		}
	}

	return false
}
