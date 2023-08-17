package main

import "fmt"

func main() {
	Greet()

	func() {
		str := "Welcome to Victoria's Secret & Co. USA"
		fmt.Println(str)

	}() // Exec

	func(str string) {
		fmt.Println(str)
	}("Welcome to Victoria's Secret & Co. Europe")

	rstr := func(str string) string {
		reverse := ""
		for _, v := range str {
			reverse = string(v) + reverse
		}
		return reverse
	}("Welcome to Victoria's Secret & Co.")
	fmt.Println("Reverse: ", rstr)

	f1 := func(str string) string { // assigning a function to a variable
		reverse := ""
		for _, v := range str {
			reverse = string(v) + reverse
		}
		return reverse
	}
	rstr = f1("Welcome to Victoria's Secret & Co.")
	fmt.Println("Reverse: ", rstr)

}

func add(a, b int) int {
	c := a + b
	return c
}

var f1 func(string) string

func Greet() {
	fmt.Println("Welcome to Victoria's Secret & Co. Bangalore")
}
