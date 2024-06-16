package main

import (
	"fmt"
	"slices"
)

func main() {
	red := []int{5, 8, 1, 3, 4}
	blue := []int{6, 9, 2, 4, 5}
	fmt.Println(classPhoto(red, blue))
	red = []int{5, 8, 1, 3, 4}
	blue = []int{9, 5, 2, 4, 5}
	fmt.Println(classPhoto(red, blue))
}

// O(n log n) Time | O(1) Space
func classPhoto(redShirtsHeights, blueShirtsHeights []int) bool {
	if len(redShirtsHeights) != len(blueShirtsHeights) {
		return false
	}
	//sort in place
	slices.Sort(blueShirtsHeights)
	slices.Sort(redShirtsHeights)

	//decide which shirt goes to the back
	var shirtToBack string
	length := len(blueShirtsHeights)
	if redShirtsHeights[length-1] > blueShirtsHeights[length-1] {
		shirtToBack = "red"
	} else {
		shirtToBack = "blue"
	}
	//loop from largest to smallest
	for i := len(redShirtsHeights) - 1; i >= 0; i-- {
		redShirtHeight := redShirtsHeights[i]
		blueShirtHeight := blueShirtsHeights[i]

		if shirtToBack == "red" {
			//red must be larger than blue if this is not true we return false
			//because no red shirt can be smaller or equal height to blue shirt
			if redShirtHeight <= blueShirtHeight {
				return false
			}

		} else {
			//blue must be larger than red
			if blueShirtHeight <= redShirtHeight {
				return false
			}
		}
	}
	return true
}
