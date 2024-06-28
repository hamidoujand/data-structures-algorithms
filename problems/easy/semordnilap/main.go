package main

import "fmt"

func main() {
	ss := []string{"diaper", "abc", "test", "cba", "repaid"}
	fmt.Println(semordnilap(ss))
}

// O(m * n) Time | O(n * m) Space
func semordnilap(ss []string) [][]string {
	words := make(map[string]struct{})
	var pairs [][]string

	//put every word inside of the set
	for _, s := range ss {
		words[s] = struct{}{}
	}

	//range over ss
	for _, s := range ss { // O(m)
		//make it string reverse
		bs := make([]byte, 0, len(s)) // O(n)
		for i := len(s) - 1; i >= 0; i-- {
			bs = append(bs, s[i])
		}
		reversed := string(bs)                             //also does cost O(n)
		if _, ok := words[reversed]; ok && reversed != s { //we do not want palindrome
			pairs = append(pairs, []string{s, reversed})
			//remove both
			delete(words, reversed)
			delete(words, s)
		}
	}
	return pairs
}
