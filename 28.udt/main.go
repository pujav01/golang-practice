package main

import "fmt"

func main() {
	r1 := Rect{L: 100.5, B: 170.7}
	a1 := AreaofRect(r1)
	fmt.Println("Area of rectangle r1:", a1)

	a2 := r1.Area()
	fmt.Println("Area of rectangle r1:", a2)
}

type Rect struct {
	L, B float32
	A    float32
}

type Square struct {
	S float32
}

func (r Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

// it is a function
func AreaofRect(r Rect) float32 {
	return r.L * r.B
}

func PerimeterofRect(r Rect) float32 {
	return 2 * (r.L + r.B)
}
