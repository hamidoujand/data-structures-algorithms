package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	a := []int{-1, 5, 10, 20, 28, 3}
	b := []int{26, 134, 135, 15, 17}
	fmt.Println(smallestDifference(a, b))
}

// O(n log n + m log m) Time | O(1) Space
// non-empty slices
func smallestDifference(s1, s2 []int) []int {
	result := make([]int, 2)
	//sort them
	slices.Sort(s1) // O(n log n)
	slices.Sort(s2) // O(m log m)

	s1Pointer := 0
	s2Pointer := 0

	candidateDifference := math.Inf(1)

	for s1Pointer < len(s1) && s2Pointer < len(s2) {
		num1 := s1[s1Pointer]
		num2 := s2[s2Pointer]

		var current float64
		//if first number is greater then no point to go to a bigger number because that increases the difference
		//so we move other pointer to a larger number
		if num1 > num2 {
			current = float64(num1 - num2)
			s2Pointer++
		} else if num2 > num1 {
			current = float64(num2 - num1)
			s1Pointer++
		} else {
			//equal
			result = append(result, num1, num2)
			return result
		}
		//now compare them
		if candidateDifference > current {
			candidateDifference = current
			result[0] = num1
			result[1] = num2
		}
	}
	return result
}
