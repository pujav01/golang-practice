package main

import (
	"fmt"
	"math"
)

func getBigAndSmall(arr []int) (int, int) {
	if len(arr) == 0 {
		return math.MinInt, math.MaxInt
	}

	min := math.MaxInt
	max := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

func main() {
	arr := []int{13, 5, 2, -1, 9, 6, 5, 20, 100, 7}

	min, max := getBigAndSmall(arr)

	fmt.Printf("Min element: %d\n", min)
	fmt.Printf("Max element: %d\n", max)
}
