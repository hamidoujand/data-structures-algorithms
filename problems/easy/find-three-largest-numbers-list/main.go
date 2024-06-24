package main

import "fmt"

func main() {
	numbers := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(findThreeLargestNumbers(numbers))
}

// O(n) Time  | O(1) Space
func findThreeLargestNumbers(numbers []int) []int {
	threeLargestOnes := make([]*int, 3) //need to check existence of them
	mapped := make([]int, 3)
	for _, num := range numbers {
		//if the we do not have a number as our third one OR we have but it less than current number we put current num and shift previous 1 idx down
		if threeLargestOnes[2] == nil || num > *threeLargestOnes[2] {

			updateAndShift(threeLargestOnes, num, 2)

		} else if threeLargestOnes[1] == nil || num > *threeLargestOnes[1] {

			updateAndShift(threeLargestOnes, num, 1)

		} else if threeLargestOnes[0] == nil || num > *threeLargestOnes[0] {

			updateAndShift(threeLargestOnes, num, 0)

		}
	}

	for i, ptr := range threeLargestOnes {
		mapped[i] = *ptr
	}
	return mapped
}

func updateAndShift(threeLargest []*int, number int, idx int) {
	//make the idx inclusive
	for i := 0; i <= idx; i++ {
		if i == idx {
			threeLargest[i] = &number
		} else {
			//move the rest 1 step backward by putting next item into current idx
			threeLargest[i] = threeLargest[i+1]
		}
	}
}
