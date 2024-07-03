package main

import "fmt"

func main() {
	numbers := []int{5, 1, 4, 2}
	fmt.Println(arrayOfProductsV1(numbers))
	fmt.Println(arrayOfProductsV2(numbers))
}

// O(n^2) Time | O(1) Space
func arrayOfProductsV1(arr []int) []int {
	products := make([]int, len(arr))

	for i := range products {
		products[i] = 1
	}

	for i := range arr {

		currentProd := 1

		for idx, num := range arr {
			if idx == i {
				continue
			}

			currentProd *= num
		}

		products[i] = currentProd
		currentProd = 1
	}
	return products
}

// O(n) Time | O(1) Space
func arrayOfProductsV2(arr []int) []int {
	products := make([]int, len(arr))
	left := make([]int, len(arr))
	right := make([]int, len(arr))

	for i := range products {
		products[i] = 1
	}
	for i := range left {
		left[i] = 1
	}
	for i := range right {
		right[i] = 1
	}

	//fill the left and right products,

	leftRunningProd := 1
	for i := range arr {
		left[i] = leftRunningProd                  // idx 0 , not item before it so its prod is 1 and so on
		leftRunningProd = leftRunningProd * arr[i] //calculate for the next idx
	}

	rightRunningProd := 1
	//this one needs to iterate from end
	for i := len(arr) - 1; i >= 0; i-- {
		right[i] = rightRunningProd
		rightRunningProd *= arr[i]
	}

	//now for each number we have calculated prod of it we just need to prod them
	//together
	for i := range arr {

		leftResult := left[i]
		rightResult := right[i]
		products[i] = leftResult * rightResult
	}
	return products
}
