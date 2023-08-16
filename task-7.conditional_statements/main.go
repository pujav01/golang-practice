package main

import "fmt"

func main() {

	state := "UP"

	gender := "M"

	height := 107

	age := 3

	ticket(state, gender, height, age)

}

func ticket(state string, gender string, height int, age int) {

	if state == "Karnataka" && gender == "F" {

		fmt.Println("No ticket")

	} else if state == "AP" && gender == "F" && height <= 110 && age <= 5 {

		fmt.Println("No ticket")

	} else if state == "Delhi" && gender == "F" {

		fmt.Println("No ticket")

	} else if state == "UP" && gender == "F" && height <= 120 && age <= 6 {

		fmt.Println("No ticket")

	} else if state == "Karnataka" && gender == "M" && height <= 110 && age <= 5 {

		fmt.Println("No ticket")

	} else if state == "AP" && gender == "M" && height <= 110 && age <= 5 {

		fmt.Println("No ticket")

	} else if state == "Delhi" && gender == "M" && height <= 130 && age <= 7 {

		fmt.Println("No ticket")

	} else if state == "UP" && gender == "M" && height <= 120 && age <= 6 {

		fmt.Println("No ticket")

	} else {

		fmt.Println("Full ticket")

	}

}
