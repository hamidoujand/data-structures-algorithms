package main

import (
	"errors"
	"fmt"
	"math"
)

/*

Sliding Window:

useful on data structures such as "slices" and "arrays" and "strings"
and we keeping track of a "sub-set" of the data for example inside of a string
looking for the longest unique set of chars or another example is to find
the max sum of numbers inside of a sub-slice...

we are going to create a "window", it can be a "sub-array" or "sub-string"
that we can slide it around from left to right based on conditions
*/

func main() {
	fmt.Println(maxSubSliceSum([]int{-2, -6, -9, -2, -1, -8, -5, 6, 3}, 2))
}

// maxSubSliceSum takes a slice and count and will return the max sum of "count"
// sequence of numbers inside of "s"
func maxSubSliceSum(s []int, count int) (int, error) {
	if len(s) < count {
		return 0, errors.New("not enough numbers inside of slice")
	}
	sum := 0
	tempSum := 0

	//first we sum as many count says from start of slice into sum
	for i := 0; i < count; i++ {
		fmt.Println(i)
		sum += s[i]
	}

	//now we assign that into tempSum
	tempSum = sum
	//start from what we already sum together and loop from there
	for i := count; i < len(s); i++ {
		//now we take previous sum and "add" next value to it and subtract
		//the "first"
		tempSum = tempSum + s[i] - s[i-count] //add next value and subtract 1 from beginning
		sum = int(math.Max(float64(tempSum), float64(sum)))
	}

	return sum, nil
}
