package main

import (
	"fmt"
)

/*
Binary Heaps: its a tree data structure, similar to binary search trees, but has different rules.
for example in a "max binary heap", "parent node" is always greater than the children
and for "min binary heap" the opposite is true, "parent is always smaller than the children"

it's called "binary" because each "node" can have at most 2 children

NOTE: there is no "ordering" between siblings. the only rule is that the "parent" is either greater or smaller based
on the type of the "binary heap"

always insert the "left" child first then the "right" child

we use a "slice" to implement it, and for parent at index "n" its left child will be at "2n +1" and right child "2n+2"
and for any "child" at index "n", its parent is at idx "(n-1)/2"


Big O:
insertion: O(log n)
deletion: O(log n)
search: O(n) not used for searching usually.


*/

type MaxBinaryHeap struct {
	values []int
}

func (h *MaxBinaryHeap) Insert(val int) {
	//append it to the end
	h.values = append(h.values, val)
	//now bubble it up
	h.bubbleUp()
}

func (h *MaxBinaryHeap) bubbleUp() {
	//get the last element from the values
	idx := len(h.values) - 1
	val := h.values[idx]
	//now we go into a loop
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		parent := h.values[parentIdx]
		//if value is not greater or its equal, break
		if parent >= val {
			break
		}
		//swap parent and child
		h.values[parentIdx], h.values[idx] = h.values[idx], h.values[parentIdx]
		//now calculate the next "parentIdx"
		idx = parentIdx
	}
	//the value is in correct spot
}

func (h *MaxBinaryHeap) ExtractMax() int {

	//swap first element with the last one first
	h.values[0], h.values[len(h.values)-1] = h.values[len(h.values)-1], h.values[0]
	//now the max is at the end we extract it
	max := h.values[len(h.values)-1]
	//shrink down the slice
	h.values = h.values[:len(h.values)-1]
	//if there are more than 1 element inside of the slice sink down.
	if len(h.values) > 1 {
		//now we sink the new root to its correct place.
		h.sink()
	}

	return max
}

func (h *MaxBinaryHeap) sink() {
	idx := 0
	length := len(h.values)
	element := h.values[0]

	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2

		var leftChild int
		var rightChild int
		var swapIdx *int //nil by default
		//check for idex to be in bound
		if leftChildIdx < length {
			leftChild = h.values[leftChildIdx]
			if leftChild > element {
				swapIdx = &leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = h.values[rightChildIdx]
			//swap if there is no "swapIdx" set, means "left" is smaller OR swap if "swapIdx" is set but the "rightChild"
			//is bigger than left child so we change the "swapIdx" to the idx of right Child
			if (swapIdx == nil && rightChild > element) || (swapIdx != nil && rightChild > leftChild) {
				swapIdx = &rightChildIdx
			}
		}

		if swapIdx == nil {
			break
		}
		//otherwise do the swap
		h.values[idx], h.values[*swapIdx] = h.values[*swapIdx], h.values[idx]
		//anc change the "idx" to be the new swapped idx
		idx = *swapIdx
	}
}

func main() {
	h := MaxBinaryHeap{
		values: make([]int, 0, 10),
	}
	h.Insert(41)
	h.Insert(39)
	h.Insert(33)
	h.Insert(18)
	h.Insert(27)
	h.Insert(12)
	h.Insert(55)
	fmt.Println(h.ExtractMax())
	fmt.Println(h.values)
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
	fmt.Println(h.ExtractMax())
}
