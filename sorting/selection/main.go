package main

import "fmt"

/*
Selection Sort: we select the min value at each loop and put it at the beginning of slice

Big O(N^2)

compare to "bubble sort" it has 1 swap per loop.

*/

func main() {
	items := []int{100, 20, -10, 1, 0, 20, 200, 2, 102, 12}
	selectionSort(items)
	fmt.Println(items)
}

func selectionSort(items []int) {
	for i := 0; i < len(items); i++ {
		//in each loop we sort the start of the slice, so we need to skip the already sorted portion "j:=i+1"
		for j := i + 1; j < len(items); j++ {
			min := items[i]               //current "min" is from first loop
			if items[j] < min && i != j { // "next candidate min" is from "nested loop"
				// "i!=j" is an optimization so if the "min" that we initially begin with is the min and nothing is greater
				// we do not swap. its for optimization
				items[i], items[j] = items[j], items[i] //so in here "i" means "first loop" and "j" means "nested loop"
			}
		}
	}
}
