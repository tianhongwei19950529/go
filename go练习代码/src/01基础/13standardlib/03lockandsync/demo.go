package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a = 0
	var b sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			b.Lock()
			a += 1
			fmt.Println(a, "in lock")
			b.Unlock()
			fmt.Println(a)
		}(i + 1)
		go func(idx int) {
			b.Lock()
			a -= 1
			b.Unlock()
		}(i - 1)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
