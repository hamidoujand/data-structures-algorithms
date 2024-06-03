package main

import "fmt"

func main() {
	cache := make(map[int]int)
	fmt.Println(fib(20))
	fmt.Println(fibMemo(20, cache))
	fmt.Println(fibTab(20))
}

// Big O : O(2^n) this is one of the worst performances, exponential, because we are
// repeating the calculation of the sub-problems over and over again.
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// Big O: O(n) because only 1 time we calculate "fib()" for each number after that we use the cached values
// since uses recursion, we have the risk of "stack overflow"
func fibMemo(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	if n <= 2 {
		return 1
	}

	res := fibMemo(n-1, cache) + fibMemo(n-2, cache)
	cache[n] = res
	return res
}

// Big O: O(n) since its a loop based on input n.  smaller space complexity, no risk for stack overflow
func fibTab(n int) int {
	if n <= 2 {
		return 1
	}
	length := n + 3
	fibs := make([]int, length)
	fibs[0] = 0
	fibs[1] = 1
	fibs[2] = 1
	for i := 3; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[n]
}
