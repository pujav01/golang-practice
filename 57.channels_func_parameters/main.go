package main

import "fmt"

func main() {
	ch := make(chan int)
	notify := make(chan struct{})
	go receiver(ch, notify)
	go sender(ch)
	<-notify
}

func sender(ch chan<- int) {
	for i := 0; i <= 100; i++ {
		ch <- i * i
	}

}

func receiver(ch <-chan int, notify chan struct{}) {
	for i := 1; i <= 100; i++ {
		fmt.Println("received", <-ch)
	}
	notify <- struct{}{}

}
