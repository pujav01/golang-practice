package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel
	ch <- 100
	ch <- 200
	// ch <- 300 // blocked because 2 is the length
	fmt.Println(<-ch, <-ch)
	ch <- 300
	ch <- 400
	fmt.Println(<-ch, <-ch)

}
