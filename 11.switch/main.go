package main

func main() {
	switch day := 4; day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	char := 'A'
	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is a uppercase char")

	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is a lowercase char")

	}
}
