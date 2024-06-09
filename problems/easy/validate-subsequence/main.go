package main

func main() {

}

// O(n) Time | O(1) Space
func validateSubsequence(seq []int, sub []int) bool {
	seqPointer := 0
	subPointer := 0
	//check both pointers to not pass the corresponding slice bound
	for seqPointer < len(seq) && subPointer < len(sub) {
		if seq[seqPointer] == sub[subPointer] { //if they match, move the "sub" by one
			subPointer++
		}
		//no matter what, always move the seqPointer forward
		seqPointer++
	}
	return seqPointer == len(sub)
}

// O(n) Time | O(1) Space
func validateSubsequenceV2(seq []int, sub []int) bool {
	subPointer := 0
	for _, num := range seq {
		if subPointer == len(sub) {
			break
		}
		if sub[subPointer] == num {
			subPointer++
		}
	}
	return subPointer == len(sub)
}
