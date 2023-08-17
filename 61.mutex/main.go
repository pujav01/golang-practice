package main

import (
	"fmt"
	"runtime"
)

var num int

func main() {

	for i := 0; i <= 1000; i++ {
		go Incr()
	}
	fmt.Println("Final value of num:", num)
	runtime.Goexit()
}

func Incr() {
	num++
	fmt.Println(num)
}
