package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	txt := "AAAAAAAAAAAAABBCCCCDD"
	fmt.Println(runLineEncoding(txt))
}

// O(n) Time | O(n) Space
// non-empty string, ASCII
func runLineEncoding(str string) string {
	encodedChars := make([]string, 0, len(str)) //to avoid "O(n)" in case of full capacity of backing array
	runCount := 1                               //since we are dealing with non-empty string, so always we have 1 of some char

	for i := 1; i < len(str); i++ {
		//we compare current char with one before it, if they are the same we increment the runCount
		currentChar := str[i]
		oneBefore := str[i-1]

		if currentChar != oneBefore || runCount == 9 {
			//encode it
			encodedChars = append(encodedChars, strconv.Itoa(runCount), string(oneBefore))
			runCount = 0 //reset it to 0
		}
		//here we inc it no matter what
		runCount++
	}

	//when the loop finishes, the last char never encoded, since the loop never ran to get to the "if/else"
	//so we need to manually add the last one
	encodedChars = append(encodedChars, strconv.Itoa(runCount), string(str[len(str)-1]))
	var builder strings.Builder
	for _, ch := range encodedChars {
		builder.WriteString(ch)
	}
	return builder.String()
}
