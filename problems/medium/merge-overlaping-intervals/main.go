package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	intervals := [][]int{
		{1, 2},
		{3, 5},
		{4, 7},
		{6, 8},
		{9, 10},
	}

	fmt.Println(mergeOverlappingIntervals(intervals))
}

// O(n log n) Time | O(n) Space
func mergeOverlappingIntervals(intervals [][]int) [][]int {
	//sort it based on their start time
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	mergedIntervals := make([][]int, 0, len(intervals)) //worst case scenario

	currentInterval := intervals[0]

	//append the first one so we can compare the rest to it when merging
	mergedIntervals = append(mergedIntervals, currentInterval)

	for _, interval := range intervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := interval[0], interval[1]

		//they overlap each other if "currentIntervalEnd >= nextIntervalStart"
		//otherwise they do not overlap each other

		if currentIntervalEnd >= nextIntervalStart {
			//overlap, merge it: we change the "currentIntervalEnd" to be the "max" between "currentIntervalEnd" and "nextIntervalEnd"
			currentInterval[1] = int(math.Max(float64(currentIntervalEnd), float64(nextIntervalEnd)))
		} else {
			//do not overlap
			//add it to the merged list as it is
			mergedIntervals = append(mergedIntervals, interval)
		}

		//change the current interval to be last item inside of "mergedIntervals"
		currentInterval = mergedIntervals[len(mergedIntervals)-1]
	}

	return mergedIntervals
}
