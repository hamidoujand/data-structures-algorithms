package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralTraverse(matrix))
	fmt.Println(spiralTraverseRecursive(matrix))
}

// O(n) Time | O(n) Space , n is total numbers of elements in entire matrix
func spiralTraverse(matrix [][]int) []int {
	var result []int

	//define bounds
	startRow := 0             //first arr
	endRow := len(matrix) - 1 //last arr

	startCol := 0 //first item inside of first row
	endCol := len(matrix[0]) - 1

	//as long as there are things to visit
	for startRow <= endRow && startCol <= endCol {
		//iterate over top bound
		for col := startCol; col <= endCol; col++ {
			result = append(result, matrix[startRow][col])
		}

		//iterate over right bound
		// "startRow +1" since we do now want to count the same item at corner twice
		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, matrix[row][endCol])
		}

		//iterate over bottom bound backwards
		for col := endCol - 1; col >= startCol; col-- {
			result = append(result, matrix[endRow][col])
		}

		//iterate over left bound upwards
		for row := endRow - 1; row >= startRow+1; row-- {
			result = append(result, matrix[row][startCol])
		}

		//now move inwards
		startRow++
		endRow--
		startCol++
		endCol--
	}

	return result
}

// O(n) Time | O(n) Space
func spiralTraverseRecursive(matrix [][]int) []int {
	var result []int
	// define bounds
	startRow := 0             //first arr
	endRow := len(matrix) - 1 //last arr

	startCol := 0 //first item inside of first row
	endCol := len(matrix[0]) - 1

	recursive(matrix, startRow, endRow, startCol, endCol, &result)
	return result
}

func recursive(matrix [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	//base case
	if startRow > endRow || startCol > endCol {
		return
	}

	//scan the outer bounds
	for col := startCol; col <= endCol; col++ {

		*result = append(*result, matrix[startRow][col])
	}

	//iterate over right bound
	// "startRow +1" since we do now want to count the same item at corner twice
	for row := startRow + 1; row <= endRow; row++ {
		*result = append(*result, matrix[row][endCol])
	}

	//iterate over bottom bound backwards
	for col := endCol - 1; col >= startCol; col-- {
		*result = append(*result, matrix[endRow][col])
	}

	//iterate over left bound upwards
	for row := endRow - 1; row >= startRow+1; row-- {
		*result = append(*result, matrix[row][startCol])
	}

	recursive(matrix, startRow+1, endRow-1, startCol+1, endCol-1, result)
}
