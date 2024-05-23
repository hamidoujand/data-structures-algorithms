package main

import "fmt"

/*
Insertion Sort: its going to treat one portion of the slice as sorted and then it
takes elements and rearrange them in that sorted section and create the slice gradually.

Big O(N^2)

if the data almost already sorted, its good, another scenario that this algorithm shines is when data coming online
and needs sorting. because it keeps one side of slice sorted and inserts 1 item at the time in right place
*/

func main() {
	items := []int{10, 2, 3, 7, 1}
	insertionSort(items)
	fmt.Println(items)
}

func insertionSort(items []int) {
	for i := 1; i < len(items); i++ {
		//we select each item inside of the slice after index 0, since that is going to be our sorted portion
		current := items[i]
		//now we iterate over numbers before the current value if the current is larger than the "previous numbers"
		var j = i - 1 //we need the "j" after the loop
		for ; j >= 0 && items[j] > current; j-- {
			//we move "larger one" 1 index forward
			items[j+1] = items[j]
		}
		//and when we moved all the larger ones forward now its time to insert the "current"
		//we want to insert it at "j+1" because when we "left" the loop "j" must be at item that
		//is not greater than the current that we exited, so we need to insert at one after that "j+1"

		items[j+1] = current
	}
}
