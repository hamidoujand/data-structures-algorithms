package main

import "fmt"

func main() {
	str := "werewolf and wolf and elf "
	fmt.Println(naiveSubStringSearch(str, "lf"))
}

func naiveSubStringSearch(str, sub string) int {
	var count int
	//loop the string

	for i := 0; i < len(str); i++ {
		//for each char we start a loop from sub we are looking for
		for j := 0; j < len(sub); j++ {
			//we break out of this loop an move to next iteration of the outer loop
			//as soon as the next char from string is not the same as sub
			if sub[j] != str[i+j] {
				break
			}
			//we have found 1 match because "j" was able to do 1 iteration over "sub"
			if j == len(sub)-1 {
				count++
			}
		}
		//otherwise we end up in here which means we have a match

	}
	return count
}
