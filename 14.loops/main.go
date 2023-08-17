package main

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
			continue
		}
		print(i, " ")
	}

	num := 10

	i := 2 // initializing i

	isPrime := true
	var count uint8 = 0

	for ; i <= num%2; i++ {
		if num/i == 0 {
			println(num, "num is not prime")
			isPrime = false
			count++
			break
		}
	}
	if isPrime == true {
		println("Prime")
	} else {
		println("Not Prime")
	}

}
