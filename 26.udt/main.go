package main

import "fmt"

func main() {
	var p1 Person
	p1.id = 101
	p1.name = "victoria"
	p1.email = "abc@victoria.com"
	p1.mobile = "9891892891"

	p2 := Person{id: 102, name: "vs", email: "vs@victoria.com", mobile: "7643271234"} // short-hand declaartion
	/*
		p3 := Person()
		p3.id = 103
		p3.name = "victoria1"
		p3.email = "abc@victoria1.com"
		p3.mobile = "98918928676"
	*/
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

}

type Person struct {
	id     int
	name   string
	email  string
	mobile string
}
