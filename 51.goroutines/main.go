package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 1; i <= 1000; i++ {
			go IsEven(i)

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
