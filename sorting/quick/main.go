package main

import "fmt"

/*
Quick Sort: we split slice into 1 or 0 element slices which any slice with 0 or 1
element inside of it is already sorted.
then we are going to select a number inside of the slices and move all numbers smaller to left and all
number greater to the right side, note that we are not sorting we just move the numbers around compare to that
selected number, this way we know that, that selected number is is correct spot

Best Time : Big O(n Log n)
Worst: Big O(N^2) because it depends on the "pivot" index, if the data already sorted then your pivot wont give you a portion
of the slice it will be one slice but 1 item less.

*/

func main() {
	s := []int{100, -3, 2, 4, 6, 9, 1, 2, 5, 3, 23}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}

func pivot(s []int, lower int, upper int) int {
	selected := s[lower] //"lower bound" is the selected element
	swapIndex := lower   //what ever the "lower bound" is, its our swapIndex start position

	//we loop from "lower + 1" till "upper included" and compare them to "selected"
	for i := lower + 1; i <= upper; i++ {
		if s[i] <= selected { //if value is smaller than "select" we first move index by one and swap the "smaller value"
			//with whatever is currently at "swapIndex"

			//then we increase our swapIndex
			swapIndex++
			//we swap "i" with "swapIndex"
			s[i], s[swapIndex] = s[swapIndex], s[i]
		}
	}
	//at the end we know the correct index of "selected" and we swap to it
	s[lower], s[swapIndex] = s[swapIndex], s[lower]
	return swapIndex
}

func quickSort(s []int, left int, right int) {
	if left < right {
		//because length of slice will not change, so we can not use "length" as base case.
		//but the only input that is changing is going to be "start" and "end" so
		//when "start < end"
		pivotIdx := pivot(s, left, right) //we do not want to hardcode "0" and "len(s)-1", since we call this fn recursively
		//call quick sort on the "left" and "right"
		//left
		quickSort(s, left, pivotIdx-1)
		//right
		quickSort(s, pivotIdx+1, right)
	}
}
