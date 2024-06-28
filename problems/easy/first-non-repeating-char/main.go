package main

import "fmt"

func main() {
	txt := "abcdcaf"
	fmt.Println(firstNonRepeatingCharV1(txt))
	fmt.Println(firstNonRepeatingCharV2(txt))
}

// ASCII string
// O(n^2) Time | O(1) Space
func firstNonRepeatingCharV1(str string) int {
	for i := range str {
		foundDuplicate := false
		for j := range str {
			if i == j {
				continue
			}

			if str[i] == str[j] {
				foundDuplicate = true
			}
		}

		if !foundDuplicate {
			return i
		}

	}
	return -1
}

// O(n) Time | O(1) Space, if dealing with finite numbers of chars
func firstNonRepeatingCharV2(str string) int {
	charsCount := make(map[rune]int, 26) //since dealing with lowercase chars only , O(1) space
	for _, ch := range str {
		charsCount[ch]++
	}

	//now we iterate over the str and check to find the first char with count of 1
	for idx, ch := range str {
		if v := charsCount[ch]; v == 1 {
			return idx
		}
	}
	return -1
}
