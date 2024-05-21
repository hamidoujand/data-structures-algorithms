package main

import "fmt"

/*
Multiple-Pointers: used on linear structures such as "string" "slice" "arrays" "linked lists"

in this pattern we are usually going to create pointer or values that correspond
to indexes and then based on the solution and conditions, we move those pointer
towards "start" "end" or "middle" or "each other"


*/

func main() {
	// s := []int{-4, -3, -2, -1, 0, 1, 2, 3, 10}
	// s2 := []int{1, 2, 3}
	// fmt.Println(sumZero(s))
	// fmt.Println(sumZero(s2))
	s := []int{1, 1, 1, 1, 2, 3, 4, 4, 4, 4, 5}
	fmt.Println(countUniqueValues(s))
	s = []int{1, 1, 2}
	fmt.Println(countUniqueValues(s))
}

// sumZero takes a slice of sorted ints and return the first pair that would result
// to zero.
func sumZero(s []int) []int {
	//because the slice is sorted from "smallest----->largest" so to get a sum
	//0 we need to add "smallest + largest" to see if we can get it if the result
	//is "+" positive then we move upper pointer down if results are "-" negative
	//we move lower pointer up till get 0 or the pointers reach each other

	left := 0
	right := len(s) - 1
	var result []int
loop:
	for left < right {
		sum := s[left] + s[right]
		switch {
		case sum == 0:
			result = append(result, s[left], s[right])
			break loop
		case sum > 0:
			right--
		case sum < 0:
			left++
		}
	}
	return result
}

// countUniqueValues will take a sorted slice of ints and it will count the number
// of unique items in there. you are allowed to alter the slice.
func countUniqueValues(s []int) int {
	if len(s) == 0 {
		return 0
	}
	counter := 0
	//scout starts at 1, and goes till end of the slice, and if the value that
	//scout pointing is not the same to the value that counter pointing we add
	//to counter and we put the value in the new position
	for scout := 1; scout < len(s); scout++ {
		current := s[scout]
		if current != s[counter] {
			counter++
			//we DO not swap them, we just put current into counter position
			s[counter] = current
		}
	}
	return counter + 1 //since "counter" correspond to an index
}
