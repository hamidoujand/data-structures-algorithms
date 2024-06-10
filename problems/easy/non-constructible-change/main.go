package main

import (
	"fmt"
	"slices"
)

/*
Given array of coins, find the minimum amount of change you can make with those coins and create new values (money)
*/

func main() {
	coins := []int{5, 7, 1, 1, 3, 22}
	fmt.Println(nonConstructibleChanges(coins))
}

// O(n log n) Time | O(1) Space
func nonConstructibleChanges(coins []int) int {
	//first we sort it
	slices.Sort(coins) // "n log n" comes from here.
	currentChangeWeCanMake := 0
	for _, coin := range coins {

		if coin > currentChangeWeCanMake+1 {
			//we have a large coin that is larger than amount of changes we already can make
			//so we do not add this one and return how many change prior to this we could do
			return currentChangeWeCanMake + 1
		}
		currentChangeWeCanMake += coin
	}
	return currentChangeWeCanMake + 1
}
