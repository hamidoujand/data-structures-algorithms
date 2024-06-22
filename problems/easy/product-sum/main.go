package main

import "fmt"

func main() {
	arr := []any{
		5,
		2,
		[]int{7, -1},
		3,
		[]any{6, []int{-13, 8}, 4},
	}
	fmt.Println(productSum(arr))
}

// O(n) Time (including  items inside of sub-arrays) | O(d) max depth of sub-arrays
func productSum(array []any) int {
	m := 1
	return recursive(array, m)
}

func recursive(array []any, multiplier int) int {
	sum := 0
	for _, item := range array { //base case is this range loop
		//check to see if item is a number of slice of int
		switch t := item.(type) {
		case int:
			sum += t
		case []any:
			sum += recursive(t, multiplier+1)
		}
	}
	return sum * multiplier
}
