package main

import "fmt"

func main() {
	// using fmt.Println
	fmt.Println("Hello World")

	// using fmt.Print
	fmt.Print("Hello")

	// using fmt.Printf
	a, b := 10, 10.5
	fmt.Printf("Type of a %T and Type of b %T", a, b)
}
