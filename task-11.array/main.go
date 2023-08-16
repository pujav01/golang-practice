package main

import "fmt"

func transposeMatrix(arr [][]int) [][]int {
	rows := len(arr)
	cols := len(arr)

	transpose := make([][]int, cols)
	for i := 0; i < cols; i++ {
		transpose[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = arr[i][j]
		}
	}
	return transpose
}

func main() {
	array2d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	transpose := transposeMatrix(array2d)

	fmt.Println("Original array:")
	for _, row := range array2d {
		fmt.Println(row)
	}

	fmt.Println("Transposed array:")
	for _, row := range transpose {
		fmt.Println(row)
	}
}
