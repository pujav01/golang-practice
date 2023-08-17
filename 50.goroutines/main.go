package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	go fmt.Println("Welcome to Victoria's Secret & Co.")
	go func() {
		for i := 1; i <= 1000; i++ {
			num := rand.Intn(1000)
			go IsEven(num)
		}
	}()

	runtime.Goexit()
}

func IsEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "Is even")
	} else {
		fmt.Println(num, "Is odd")
	}
}
