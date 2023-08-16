package main

import "fmt"

var myvar1 = 100

func main() {
	var myvar1 = 50

	fmt.Println("myvar1:", myvar1) // prints the value in the nearest scope i.e 50
}
