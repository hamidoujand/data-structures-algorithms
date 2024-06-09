package main

import "slices"

/*
Two Number Sum
*/

func main() {

}

// O(n^2) Time | O(1) Space
func twoNumberSumNaive(s []int, target int) []int {
	var result []int
	for i := 0; i < len(s); i++ {
		firstNum := s[i]
		for j := i + 1; j < len(s); j++ {
			secondNum := s[j]
			sum := firstNum + secondNum
			if sum == target {
				result = append(result, firstNum, secondNum)
				return result
			}
		}
	}
	return result
}

// O(n) Time | O(n) Space
func twoNumberSumHash(s []int, target int) []int {
	hash := make(map[int]struct{}, len(s))
	var result []int

	for i := 0; i < len(s); i++ {
		currentNum := s[i]
		y := target - currentNum
		if _, ok := hash[y]; ok {
			result = append(result, currentNum, y)
		} else {
			hash[y] = struct{}{}
		}
	}
	return result
}

// slice must be sorted or you will be allowed to sort it
// O(n log n) Time | O(1) Space
func twoNumberSumTwoPointers(s []int, target int) []int {
	var result []int
	slices.Sort(s)
	leftPointer := 0
	rightPointer := len(s) - 1
	for leftPointer < rightPointer { //as long as the left did not overlap the right one
		firstNum := s[leftPointer]
		secondNum := s[rightPointer]

		candidateSum := firstNum + secondNum
		if candidateSum == target {
			result = append(result, firstNum, secondNum)
			return result
		} else if candidateSum > target {
			//since it's sorted, we need to move away from the end
			rightPointer--
		} else {
			leftPointer++
		}

	}
	return result
}
