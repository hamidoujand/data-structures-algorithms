package main

import "fmt"

/*
Bubble Sort: not that efficient and commonly used
in this sorting algo, we are going to compare two values and if the first one
is larger then we swap them till the largest one bubbles to the end of slice

if the data is already sorted or mostly sorted, the logic runs no matter what
and iterate over those sorted we can  optimize this by checking if we swapped
last time or not if we did not swap then that means data is sorted

Big O(N^2)

if the data already sorted and we use "swap" to make sure we are not wasting loops
on already sorted data then its good algorithm
*/

func main() {
	items := []int{37, 49, 27, 8, -1, 0, 2}
	bubbleSort(items)
	fmt.Println(items)
}

func bubbleSort(items []int) {

	for i := 0; i < len(items)-1; i++ {
		noSwaps := true //default we assume no swap will happen

		//in each loop there are going to be some items already sorted at the end of slice
		//so we need to ignore them
		for j := 0; j < len(items)-1-i; j++ {
			//using "j" as the index to work with, because we need to do operation per each number inside of slice.
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				noSwaps = false //we did swap
			}
		}
		//here we check
		if noSwaps {
			break
		}
	}
}
