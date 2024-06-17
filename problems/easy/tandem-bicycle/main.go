package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	blues := []int{5, 5, 3, 9, 2}
	reds := []int{3, 6, 7, 2, 1}
	fmt.Println(tandemBicycle(blues, reds, true))
}

// O(n log n) Time | O(1) Space
func tandemBicycle(redShirts, blueShirts []int, fastest bool) int {
	//sort them
	slices.Sort(redShirts)
	slices.Sort(blueShirts)
	total := 0
	if !fastest {
		//reverse one of them
		reverse(redShirts)
	}

	for i, j := 0, len(redShirts)-1; i < len(blueShirts) && j >= 0; i, j = i+1, j-1 {
		rider1 := float64(blueShirts[i])
		rider2 := float64(redShirts[j])
		max := math.Max(rider1, rider2)
		total += int(max)
	}

	return total
}

func reverse(s []int) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
