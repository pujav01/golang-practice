package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {

		go func() {
			for i := 1; i <= 100; i++ {
				ch1 <- i
			}
			close(ch1)
		}()

		ch2 <- 200
		close(ch2)
		ch3 <- 300
		close(ch3)
	}()
	ch4 := VariadicChan(ch1, ch2, ch3)
	for val := range ch4 {
		fmt.Println(val)
	}

}

func VariadicChan(chs ...chan int) chan int {
	ch := make(chan int)
	go func() {
		for _, ch := range chs {
			sum := 0
			for val := range ch {
				sum += val
			}
			ch <- sum

		}
		close(ch)
	}()

	return ch
}
