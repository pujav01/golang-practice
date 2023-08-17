package main

func main() {
	i := 1

loop:

	if i%2 == 0 {
		println(i, "is an even number")
	}
	i++
	if i <= 10 {
		goto loop
	}

}
