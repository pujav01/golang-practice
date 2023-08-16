package main

import (
	"fmt"
	"math"
)

func secondBiggest(arr []int) int {
	if len(arr) < 0 {
		return math.MinInt
	}

	largest := arr[0]
	secondLargest := math.MinInt

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] < largest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}

func secondSmallest(arr []int) int {
	if len(arr) < 0 {
		return math.MaxInt
	}

	smallest := arr[0]
	secondSmallest := math.MaxInt

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			secondSmallest = smallest
			smallest = arr[i]
		} else if arr[i] < secondSmallest && arr[i] > smallest {
			secondSmallest = arr[i]
		}
	}
	return secondSmallest
}

func main() {
	array := []int{1, 2, 5, 6, 3, 8, 10, 16}

	secondLargest := secondBiggest(array)
	secondSmallest := secondSmallest(array)

	fmt.Printf("Second biggest element in the array: %d\n", secondLargest)
	fmt.Printf("Second smallest element in the array: %d\n", secondSmallest)
}
