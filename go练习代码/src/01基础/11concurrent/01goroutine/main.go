package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done()
}

//func singe() {
//	wg.Add(1)
//	go hello()
//	fmt.Println("hello main")
//	wg.Wait()
//}

func more() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		hello(i)
	}
	fmt.Println("hello main")
	wg.Wait()
}
func main() {
	//singe()
	more()
}
