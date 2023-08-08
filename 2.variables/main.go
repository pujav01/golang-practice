package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age uint = 38
	fmt.Println("Age:", age, "Type of age:", reflect.TypeOf(age))
	fmt.Printf("\nAge: %v Type of Age %T", age, age)

	var num1 = 100
	var num2 = 25.24
	num3 := 1234
	fmt.Println(num1, num2, num3, reflect.TypeOf(num1), reflect.TypeOf(num2), reflect.TypeOf(num3))

	// Type interface
	var (
		num6 int     //0
		num7 float32 //0
		ok1  bool    //false
		str1 string  //false
	)
	fmt.Println(num6, num7, ok1, str1)

}
