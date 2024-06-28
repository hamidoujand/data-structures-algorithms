package main

import "fmt"

func main() {
	chars := "nma !eht si edo nhjo"
	doc := "john doe is the man!"
	fmt.Println(generateDocumentV1(chars, doc))
	fmt.Println(generateDocumentV2(chars, doc))
	fmt.Println(generateDocumentV3(chars, doc))
}

// O(m * (m+n)) Time, m length of docs and n length of chars | O(1) Space
func generateDocumentV1(chars, doc string) bool {
	for _, ch := range doc { // O(m)
		//we count the occurrence of each char in the doc from start to end
		//and then check it with the number of occurrence from chars to see if
		//we have enough chars
		var docFrequency int
		for _, candidate := range doc { // O(m)
			if ch == candidate {
				docFrequency++
			}
		}

		//now we loop the chars
		var charsFrequency int
		for _, target := range chars { //O(n)
			if target == ch {
				charsFrequency++
			}
		}

		if charsFrequency < docFrequency {
			return false
		}
	}
	return true
}

// O(c * (m+n)) Time: c count of unique chars, m length of docs and n length of chars | O(1) Space
func generateDocumentV2(chars, doc string) bool {
	//we keep track of all chars that we already counted, so we optimize it a little bit
	set := make(map[rune]struct{})

	for _, ch := range doc { // O(c)
		if _, ok := set[ch]; ok {
			continue
		}
		var docFrequency int
		for _, candidate := range doc { // O(m)
			if ch == candidate {
				docFrequency++
			}
		}

		//now we loop the chars
		var charsFrequency int
		for _, target := range chars { //O(n)
			if target == ch {
				charsFrequency++
			}
		}

		if charsFrequency < docFrequency {
			return false
		}

		//add the char into set
		set[ch] = struct{}{}
	}
	return true
}

// O(n + m) Time | O(c) Space, count of unique chars
func generateDocumentV3(chars, doc string) bool {
	//uses frequency counter pattern
	charsCount := make(map[rune]int)

	for _, ch := range chars {
		charsCount[ch]++
	}

	for _, ch := range doc {
		//if there are not enough of this char inside of charCount we return false
		if v, ok := charsCount[ch]; !ok || v <= 0 {
			return false
		} else {
			charsCount[ch]--
		}
	}

	return true
}
