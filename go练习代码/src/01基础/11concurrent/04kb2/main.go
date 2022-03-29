package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n)
	wg.Done()
}

func main() {
	const Num = 100000
	wg := sync.WaitGroup{}
	for i := 0; i < Num; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
}
