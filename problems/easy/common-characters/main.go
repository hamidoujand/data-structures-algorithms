package main

import (
	"fmt"
	"maps"
)

// in order a char to be a common character, it must be in all strings
func main() {
	ss := []string{"abc", "bcd", "cbaccd"}
	fmt.Println(string(commonCharactersV1(ss)))
	fmt.Println(string(commonCharactersV2(ss)))
}

// O(n * m) Time, n length of input and "m" is length of largest str | O(c) Space c total number of unique chars, after the loop that set is gone.
func commonCharactersV1(ss []string) []rune {
	characterCount := make(map[rune]int)
	for _, str := range ss {
		//create a unique SET of chars and then insert them into charCount map or inc their count
		set := make(map[rune]struct{})
		for _, ch := range str {
			if _, ok := set[ch]; !ok {
				set[ch] = struct{}{}
			}
		}

		//now we add them into our frequency counter map
		for ch := range set {
			characterCount[ch]++
		}
	}

	//now based on the length of the ss, if a char has that many appearance inside of charCount
	//that means that is a common char
	length := len(ss)
	var finalChars []rune
	for k, v := range characterCount {
		if v == length {
			finalChars = append(finalChars, k)
		}
	}

	return finalChars
}

// O(n * m) Time, n length of input and "m" is length of largest str |  O(m) Space, m length of largest string
func commonCharactersV2(ss []string) []rune {
	//find the smallest string, and add those to the unique char set
	smallestString := getSmallestString(ss)
	uniqueSet := make(map[rune]struct{})
	for _, ch := range smallestString {
		if _, ok := uniqueSet[ch]; !ok {
			uniqueSet[ch] = struct{}{}
		}
	}

	//loop the "ss" and take each string and check "chars" from "uniqueSet" to be inside of string if any char is not, remove
	//it from unique set
	for _, s := range ss {
		removeNonExistenceChars(s, uniqueSet)
	}

	var finalChars []rune
	for k := range uniqueSet {
		finalChars = append(finalChars, k)
	}
	return finalChars
}

func getSmallestString(ss []string) string {
	candidate := ss[0]
	for i := 1; i < len(ss); i++ {
		if len(ss[i]) < len(candidate) {
			candidate = ss[i]
		}
	}
	return candidate
}

func removeNonExistenceChars(str string, potentialCharSet map[rune]struct{}) {
	//first create another set of unique chars from the string itself
	unique := make(map[rune]struct{})
	for _, ch := range str {
		if _, ok := unique[ch]; !ok {
			unique[ch] = struct{}{}
		}
	}

	//now we loop the "potentialCharSet" and remove from it any char that is not inside of other strings
	//NOTE: modifying the same "map" while "iterating over it" is not going to take effect usually so we need a clone
	//of the original map and loop that one and modify the original one
	cloned := maps.Clone(potentialCharSet)
	for k := range cloned {
		if _, ok := unique[k]; !ok {
			delete(potentialCharSet, k)
		}
	}
}
