package main

/*
need to switch the columns and rows of a matrix
*/

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	transposed := transposeMatrix(matrix)
	fmt.Println(transposed)
}

// O(width * height) Time | O(width * height) Space
func transposeMatrix(matrix [][]int) [][]int {
	transposedMatrix := [][]int{}
	//we need to find the number of columns . we need to get 1 row and see how many
	//items are inside of it and that is how many columns you have
	for col := 0; col < len(matrix[0]); col++ {
		newRow := []int{}
		//now for each column we go down row by row and add values into another slice
		for row := 0; row < len(matrix); row++ {
			newRow = append(newRow, matrix[row][col])
		}
		transposedMatrix = append(transposedMatrix, newRow)
	}
	return transposedMatrix
}
