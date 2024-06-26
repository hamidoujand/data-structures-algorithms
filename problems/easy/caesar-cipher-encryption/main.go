package main

import "fmt"

/*
Rules: only lowercases chars from a-z allowed
*/
func main() {
	txt := "xyz"
	fmt.Println(caesarCipherEncryption(txt, 2))
}

// O(n) Time  | O(n) Space
func caesarCipherEncryption(input string, key int) string {
	//a= 96
	//z= 122
	//usage of % operator to keep things in bound
	newLetters := make([]byte, 0, len(input))

	newBoundedKey := key % 26 // 26 lowercase letters

	for _, ch := range input {
		newLetterCode := ch + rune(newBoundedKey)
		if newLetterCode <= 122 {
			newLetters = append(newLetters, byte(newLetterCode))
		} else {
			newLetters = append(newLetters, byte(96+newLetterCode%122))
		}
	}
	return string(newLetters)
}
