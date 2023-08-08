package main

import "fmt"

var (
	num1 int
	num2 float32
	num3 float64
)

func main() {
	var str1 string = "Hello World"
	fmt.Println(str1)

	fmt.Println("length of str1", len(str1))

	str1 = "Hello Victoria's Secret & Co" // mutating string. STRINGS ARE IMMUTABLE FROM THE RUNTIME PERPESPECTIVE, BUT THE USER CAN REASSIGN

	fmt.Println("length of str1", len(str1))

	num1 := 10000
	fmt.Println(num1)

	num1 = 20000 // mutating a int
	fmt.Println(num1)

}
