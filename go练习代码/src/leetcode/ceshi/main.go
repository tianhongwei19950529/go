package main

import (
	"fmt"
)

func fib2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		//default:
		// fmt.Println("fuck")
		//case <-time.After(3 * time.Second):
		//	println("timeout")
		//	return
		}
	}
}

func main() {
	c := make(chan int, 2)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib2(c, quit)
}