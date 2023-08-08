package main

import "fmt"

const PI float32 = 3.14

const SQUARE_OF_PI = PI * PI // this is a valid constant

var pi float32 = 3.14

// const SQUARE_OF_PI = pi*pi // not valid

const (
	DAYS    = 7
	MONTHS  = 12
	HOURS   = 24
	MINUTES = 60
	SECONDS = 60
)

func main() {

	fmt.Println(PI, SQUARE_OF_PI, pi)
}
