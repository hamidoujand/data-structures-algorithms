package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{2, 1, 5, 3, 3, 2, 4}
	fmt.Println(findFirstDuplicateV1(numbers))
	fmt.Println(findFirstDuplicateV2(numbers))
	fmt.Println(findFirstDuplicateV3(numbers))
}

// O(n^2) Time | O(1) Space
func findFirstDuplicateV1(arr []int) int {
	minSecondIdx := len(arr) //if at the end we still pointing to this idx
	//means we do not have duplicates

	for i := 0; i < len(arr); i++ {
		current := arr[i]
		for j := i + 1; j < len(arr); j++ {
			candidate := arr[j]

			if current == candidate {
				minSecondIdx = int(math.Min(float64(minSecondIdx), float64(j)))
			}
		}
	}
	if minSecondIdx == len(arr) {
		return -1
	} else {
		return arr[minSecondIdx]
	}
}

// O(n) Time | O(1) Space
func findFirstDuplicateV2(arr []int) int {
	//we need to use a "set"
	seen := make(map[int]struct{}, len(arr))

	for _, num := range arr {
		if _, ok := seen[num]; ok {
			return num
		}
		seen[num] = struct{}{}
	}
	return -1
}

// Need the numbers to be "1...n" length and uses a weird trick that requires to mutate the arr
// O(n) Time | O(1) Space
func findFirstDuplicateV3(arr []int) int {
	for i := range arr {
		num := arr[i]
		abs := int(math.Abs(float64(num)))

		idx := abs - 1 //only works if numbers are "1...n"
		if arr[idx] < 0 {
			//we already had this value and idx
			return int(math.Abs(float64(num)))
		} else {
			//mark that number at that idx with negative sign
			arr[idx] *= -1
		}

	}

	return -1
}
