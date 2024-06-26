package main

import "fmt"

func main() {
	txt := "abcdcba"
	fmt.Println(isPalindromeIterativeV1(txt))
	fmt.Println(isPalindromeIterativeV2(txt))
	fmt.Println(isPalindromeIterativeV3(txt))
	fmt.Println(isPalindromeRecursive(txt))
}

// O(n^2) Time | O(n) Space
func isPalindromeIterativeV1(str string) bool {
	reversed := ""

	//now create the string backwards
	for i := len(str) - 1; i >= 0; i-- { //O(n)
		reversed += string(str[i]) // O(n) since requires looping []byte and create a new string each time
	}

	return reversed == str
}

// O(n) Time | O(n) Space
func isPalindromeIterativeV2(str string) bool {
	reversed := make([]byte, 0, len(str))

	//now create the string backwards
	for i := len(str) - 1; i >= 0; i-- {
		reversed = append(reversed, str[i])
	}

	return string(reversed) == str
}

// O(n) Time | O(1) Space
func isPalindromeIterativeV3(str string) bool {
	leftPointer := 0
	rightPointer := len(str) - 1
	for leftPointer <= rightPointer {
		if str[leftPointer] != str[rightPointer] {
			return false
		}
		leftPointer++
		rightPointer--
	}
	return true
}

// O(n) Time | O(n) Space
func isPalindromeRecursive(str string) bool {
	if len(str) == 1 {
		//1 char is always palindrome
		return true
	}
	return str[0] == str[len(str)-1] && isPalindromeRecursive(str[1:len(str)-1]) //in each call we create a new sub-string from second to second to last
}
