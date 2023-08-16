package main

import "fmt"

func main() {
	const (
		a = 10
		b = 20
	)

	addition := a + b
	fmt.Println("Addition: ", addition)

	subtaction := a - b
	fmt.Println("Subtaction: ", subtaction)

	multiplication := a * b
	fmt.Println("multiplication: ", multiplication)

	division := a / b
	fmt.Println("division: ", division)

	modulus := a % b
	fmt.Println("modulus: ", modulus)

}
