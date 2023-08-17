package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channnel has been initiaited
	go func() {

		ch <- 100 // Assign a value to the channel
	}()
	num := <-ch // Recieve a value from the channel

	fmt.Println(num)
}
