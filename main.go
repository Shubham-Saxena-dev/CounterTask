package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 2)
	done := make(chan bool)
	add(2, 4, ch)
	multiply(2, 4, ch)
	close(ch)
	go func(ch chan int) {
		for {
			val, ok := <-ch
			if ok {
				fmt.Println(val)
			} else {
				done <- true
				break
			}
		}
	}(ch)
	<-done
	fmt.Println("Executed")
}

func multiply(a, b int, ch chan int) {
	ch <- a * b
}

func add(a, b int, ch chan int) {
	ch <- a + b
}
