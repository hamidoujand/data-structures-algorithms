package main

import "fmt"

/*
Recursion: when a process (functions) calls itself.

its another way of solving problems, we have non-recursion way which is called "iteratively"
and we have "Recursion" we know how iterative solutions work, we iterate and go over items

we use it because sometimes its clearer and cleaner than the "iteratively" solutions

in any recursive function 2 things must be present:

-"base case": condition for recursion to stop follows by a "return" to pop that fn call from call stack

-"different input": each call to the recursive function must have a different input that goes towards the "base case"

Common Problems and Pitfalls:

-"forgetting the base case or wrong base case"
-"forgetting the 'return'"
-"returning the wrong value"
-"forgetting to change the 'input' to go towards the 'base case', otherwise it will run forever"


Helper Method Recursion: in this case we have an "outer fn" and a "separate" factorial fn
that is going to be called inside of this "outer" fn we usually use it when we need
to collect some data from the recursive call.


Pure Method Recursion: the fn itself is going to be called recursively and also return some data.


*/

func main() {
	countDown(10)
	fmt.Println(sumRange(3))
	fmt.Println(factorial(4))
	result := collectOddValues([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	fmt.Println(result)
	fmt.Println(collectEvenValues([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
}

// Recursive fn, you can see that we can also use "iterative" solution to solve the problem.
func countDown(count int) {
	//Base Case
	if count <= 0 {
		fmt.Println("All Done")
		return
	}
	fmt.Println("counter:", count)
	//Different Input
	count--
	countDown(count)
}

// you can solve it with iterative solution as well
func sumRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + sumRange(num-1)
}

// you can solve it with iterative solution as well
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// collectOddValues is an example of "helper method recursion"
func collectOddValues(s []int) []int {
	var result []int
	helper(s, &result)
	return result
}
func helper(s []int, result *[]int) {
	if len(s) == 0 {
		return
	}
	if s[0]%2 != 0 {
		*result = append(*result, s[0])
	}
	helper(s[1:], result)
}

func collectEvenValues(s []int) []int {
	var result []int //this slice will be created in each call. works because we "append" to it from other calls.
	if len(s) == 0 {
		//at the end will return a slice
		return result
	}
	if s[0]%2 == 0 {
		result = append(result, s[0])
	}
	//we wait for the slice that will be returned from "base case" and each time we append to the current stack's result slice.
	result = append(result, collectEvenValues(s[1:])...)
	return result
}
