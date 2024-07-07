package main

import "fmt"

func main() {
	seats := []int{1, 1, 1, 0, 0, 0, 1}
	fmt.Println(bestSeat(seats))
}

// 0 means empty seat, 1 means occupied, first and last seat are always 1
// O(n) Time | O(1) Space
func bestSeat(seats []int) int {
	bestSeat := -1 //default
	maxSpace := 0  //default space

	//slow pointers pattern

	left := 0
	for left < len(seats) {
		right := left + 1
		//move the right pointer till we find a occupied seat
		for right < len(seats) && seats[right] == 0 { //since it does not loop the same numbers again it does not cause "O(n^2)"
			right++
		}

		//we have left and right which both are 1, in between maybe 0 or maybe not
		availableSpace := right - left - 1 // -1 for idx start at 0 so offset

		//only seat in middle of these seats if the maxSpace here is bigger

		if availableSpace > maxSpace {
			//seat in between
			bestSeat = (left + right) / 2
			//change maxSpace as well because that is our best seat so far
			maxSpace = availableSpace
		}
		//move the "left" to the "right" idx
		left = right
	}

	return bestSeat
}
