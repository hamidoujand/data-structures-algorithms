package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	s := []int{-10, -2, -1, 0, 100, 230}
	fmt.Println(sortedSquaredSliceNaive(s))
	fmt.Println(sortedSquaredSliceTwoPointers(s))
	fmt.Println(sortedSquaredSliceTwoPointersV2(s))
}

// O(n log n) Time | O(n) Space
func sortedSquaredSliceNaive(s []int) []int {
	result := make([]int, len(s))
	for idx, num := range s {
		result[idx] = num * num
	}

	//now we sort it
	slices.Sort(result) // "n long n" comes from here
	return result
}

// O(n) Time | O(n) Space
// slice must be sorted to use this
func sortedSquaredSliceTwoPointers(s []int) []int {
	result := make([]int, len(s))
	smallerValueIdx := 0
	largerValueIdx := len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		smallerValue := s[smallerValueIdx]
		largerValue := s[largerValueIdx]

		if math.Abs(float64(smallerValue)) > math.Abs(float64(largerValue)) {
			result[i] = smallerValue * smallerValue
			smallerValueIdx++
		} else {
			result[i] = largerValue * largerValue
			largerValueIdx--
		}
	}

	return result
}

// O(n) Time | O(n) Space
func sortedSquaredSliceTwoPointersV2(s []int) []int {
	result := make([]int, len(s))
	smallerValueIdx := 0
	largerValueIdx := len(s) - 1
	idx := len(s) - 1
	for smallerValueIdx < largerValueIdx && idx >= 0 {
		smallerValue := s[smallerValueIdx]
		largerValue := s[largerValueIdx]

		if math.Abs(float64(smallerValue)) > math.Abs(float64(largerValue)) {
			result[idx] = smallerValue * smallerValue
			smallerValueIdx++
		} else {
			result[idx] = largerValue * largerValue
			largerValueIdx--
		}
		idx--
	}
	return result
}
