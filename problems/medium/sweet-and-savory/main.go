package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	dishes := []int{5, 2, -7, 30, 12, -4, -20}
	fmt.Println(sweetAndSavory(dishes, 4))
}

// find the best combination between "sweet and savor" in a way to not pass the target and be close to it
// O(n*log n) Time | O(n) Space
func sweetAndSavory(dishes []int, target int) []int {
	var sweetDishes []int
	var savorDishes []int
	results := make([]int, 2)

	for _, num := range dishes {
		if num < 0 {
			sweetDishes = append(sweetDishes, num)
		} else {
			savorDishes = append(savorDishes, num)
		}
	}

	//sort them both
	slices.Sort(savorDishes)
	//for "sweetDishes" we need sort based on the abs value
	slices.SortFunc(sweetDishes, func(a, b int) int {
		return int(math.Abs(float64(a)) - math.Abs(float64(b)))
	})

	//best difference so far
	bestDiff := math.Inf(1)

	sweetIdx := 0
	savoryIdx := 0

	for sweetIdx < len(sweetDishes) && savoryIdx < len(savorDishes) {
		currentSum := sweetDishes[sweetIdx] + savorDishes[savoryIdx]

		if currentSum <= target {
			currentDiff := float64(target - currentSum)

			if currentDiff < bestDiff {
				results[0] = sweetDishes[sweetIdx]
				results[1] = savorDishes[savoryIdx]

				bestDiff = currentDiff
			}
			//when you less than target you need to add positive numbers to get
			//close to the target
			savoryIdx++
		} else {
			sweetIdx++
		}

	}
	return results
}
