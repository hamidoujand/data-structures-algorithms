package main

import (
	"fmt"
	"math"
)

// Peak: must have at least 1 number less before and 1 number less after it to be a peak.

func main() {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	fmt.Println(longestPeak(arr))
}

// O(n) Time | O(1) Space
func longestPeak(arr []int) float64 {
	longest := 0.0

	for i := 1; i < len(arr)-1; {
		current := arr[i]
		oneBefore := arr[i-1]
		oneAfter := arr[i+1]

		isPeak := current > oneBefore && current > oneAfter
		if !isPeak {
			//we move to next value
			i++
			continue
		}

		//we found the peak, let's calculate it's length

		//since we already know "i-1" and "i+1" are part of the peak, we need to find the rest
		leftPointer := i - 2
		rightPointer := i + 2

		for leftPointer >= 0 && arr[leftPointer] < arr[leftPointer+1] {
			leftPointer--
		}

		for rightPointer < len(arr) && arr[rightPointer] < arr[rightPointer-1] {
			rightPointer++
		}

		currentPeak := rightPointer - leftPointer - 1
		longest = math.Max(float64(currentPeak), longest)
		//to find the next peak, move "i" to the end of "current peak's tail"
		i = rightPointer
	}

	return longest
}
