package main

import "fmt"

//Monotonic Array: an array that its numbers either increasing or decreasing and
//no number inside of that array breaks this direction

func main() {
	s := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println(isMonotonicV1(s))
	fmt.Println(isMonotonicV2(s))
}

// O(n) Time | O(1) Space
func isMonotonicV1(s []int) bool {
	if len(s) < 2 {
		return true
	}

	direction := s[1] - s[0] // "0" means no-direction "x>0" "upwards" "x<0" "downwards"

	for i := 2; i < len(s); i++ {
		//see if the direction changes
		if direction == 0 {
			//we need to find direction based on next value and previous one
			direction = s[i] - s[i-1]
			continue //move to next loop
		}

		if breaksDirection(direction, s[i-1], s[i]) {
			return false //not a monotonic
		}
	}
	return true
}

func breaksDirection(currentDir int, previous int, current int) bool {
	diff := current - previous
	if currentDir > 0 {
		//only breaks if "diff" is "less than 0"
		return diff < 0
	} else {
		//only breaks if "diff" is "greater than 0"
		return diff > 0
	}
}

// O(n) Time | O(1) Space
func isMonotonicV2(s []int) bool {
	isNonDecreasing := true
	isNonIncreasing := true
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			isNonDecreasing = false
		}
		if s[i] > s[i-1] {
			isNonIncreasing = false
		}
	}

	return isNonDecreasing || isNonIncreasing
}
