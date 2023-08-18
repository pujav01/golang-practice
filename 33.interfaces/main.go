package main

import "fmt"

func main() {

}

func Area(iarea IArea) { // interface is a parameter. It can be form of a dependency injection
	a := iarea.Area()
	fmt.Println("Area: ", a)
}

func Perimeter(iperimeter IPerimeter) {
	p := iperimeter.Perimeter()
	fmt.Println("Perimeter: ", p)
}

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}
