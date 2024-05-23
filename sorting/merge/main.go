package main

import "fmt"

/*
Merge Sort: it's combination of 3 things: "splitting" "sorting" "merging"

Fact: arrays of 1 or 0 elements are already sorted.

its going to decompose arrays to smaller arrays of 1 or 0 elements (divide and conquer)

Time Big O(n Log n) because "Divide and Conquer is 'Long n' " and "we loop" "n" time to merge them

Space Big O(n)

*/

func main() {
	s := []int{10, 24, 76, 73, 72, 1, 9}
	s2 := mergeSort(s)
	fmt.Println(s2)
}

func mergeSort(s []int) []int {
	//base case
	if len(s) <= 1 {
		return s
	}
	//call mergeSort on the halves
	mid := len(s) / 2
	left := mergeSort(s[:mid])
	right := mergeSort(s[mid:])

	return merge(left, right)
}

func merge(s1, s2 []int) []int {
	result := make([]int, 0, len(s1)+len(s2))

	i := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++

		}
	}
	// add the remainder
	for i := i; i < len(s1); i++ {
		result = append(result, s1[i])
	}
	for j := j; j < len(s2); j++ {
		result = append(result, s2[j])
	}

	return result
}

/*

Structure of for..loop

for <init_runs_1_time_before_loop>; <condition_will_be_check_before_each_iteration>; <steps_runs_at_the_end_of_loop> {

}

*/
