package main

import "fmt"

func arraySum(arr [3][3]int) ([]int, []int) {
	rows := len(arr)
	cols := len(arr[0])

	rowSums := make([]int, rows)
	colSums := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rowSums[i] += arr[i][j]
			colSums[i] += arr[i][j]
		}
	}

	return rowSums, colSums

}

func main() {
	array2d := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rowSums, colSums := arraySum(array2d)

	fmt.Println("Row Sum:")
	for i, sum := range rowSums {
		fmt.Printf("Row %d Sum: %d\n", i+1, sum)
	}
	fmt.Println("Column Sum:")
	for j, sum := range colSums {
		fmt.Printf("Column %d: %d\n", j+1, sum)
	}
}
