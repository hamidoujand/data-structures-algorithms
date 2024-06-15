package main

import (
	"fmt"
	"slices"
)

func main() {
	queries := []int{3, 2, 1, 2, 6}
	fmt.Println(minWaitingTime(queries))
}

// O(n log n) Time | O(1) Space
func minWaitingTime(queries []int) int {
	//1-sort it
	slices.Sort(queries)

	totalWaitingTime := 0

	for idx, duration := range queries {
		howManyLeft := len(queries) - (idx + 1)
		howMuchWait := howManyLeft * duration //all queries after it must wait at least this duration
		totalWaitingTime += howMuchWait
	}
	return totalWaitingTime
}
