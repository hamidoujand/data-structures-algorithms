package main

import "fmt"

/*

usually is going to turn "O(n^2) into  O(n)"

usually used when you need to compare your inputs and compare them, like if they
are anagram of each other or how much a char is repeated inside of string ...

we create a map and use it to break down contents of strings or slices and store
them inside of that map either for checking their existence or how often they appear
or any other profiling around them and solve related problem to them
*/

func main() {
	// s1 := []int{1, 2, 3}
	// s2 := []int{1, 4, 9}
	// fmt.Println(same(s1, s2))

	s1 := "anagram"
	s2 := "nagaram"
	fmt.Println(validAnagram(s1, s2))
}

//Example

func same(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	set1 := make(map[int]struct{}, len(s1))

	for _, item := range s1 {
		set1[item] = struct{}{}
	}

	set2 := make(map[int]struct{}, len(s2))

	for _, item := range s2 {
		set2[item] = struct{}{}
	}

	for item := range set1 {
		if _, ok := set2[item*item]; !ok {
			return false
		}
	}
	return true
}

func validAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	set := make(map[rune]int, len(s1))

	for _, ch := range s1 {
		set[ch]++
	}
	//now we range the second string, and we need to have that many runes inside
	//of the set to have an anagram
	for _, ch := range s2 {
		freq, ok := set[ch]
		switch {
		case !ok:
			// rune is not inside of "set"
			return false
		case freq == 0:
			// we already exhaust the rune appearance
			return false
		default:
			//reduce counter
			set[ch]--
		}
	}
	return true
}
