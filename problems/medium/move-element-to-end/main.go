package main

import "fmt"

func main() {
	s := []int{2, 1, 2, 2, 2, 3, 4, 2}
	moveElementToTheEnd(s, 2)
	fmt.Println(s)
}

func moveElementToTheEnd(s []int, target int) {
	left := 0
	right := len(s) - 1
	for left < right { // when "left" reaches "right" then the entire slice has been traversed
		//another loop to find the correct spot for next item to be placed
		//make sure your left and right not overlap in here
		for left < right && s[right] == target {
			right--
		}

		//now "right" is in position that the next matched number can be swapped with

		//now check the "left" pointer's value to see if it is candidate to be swapped
		if s[left] == target {
			//swap it
			s[left], s[right] = s[right], s[left]
		}
		//inc "left"
		left++
	}
}
