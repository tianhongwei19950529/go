package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//计算一个64位随机数的各位的和
func randNumber(x int64) int64 {
	var sum int64 = 0
	for x > 0 {
		a := x % 10
		x = x / 10
		sum += a
	}
	return sum
}

// 生成int64的随机数放入通道ch1中
func createRand(ch1 chan<- int64) {
	for {
		int63 := rand.Int63()
		ch1 <- int63
		time.Sleep(1)
	}
}

//从通道ch1读取数据，然后计算各个位数之和存入ch2中
func readRand(ch1 <-chan int64, ch2 chan<- int64) {
	for {
		value := <-ch1
		number := randNumber(value)
		ch2 <- number
		fmt.Println(value, number)
	}

}
func main() {
	var jobChan = make(chan int64, 100)
	var resultChan = make(chan int64, 100)
	wg.Add(25)
	go createRand(jobChan)

	for i := 0; i < 24; i++ {
		go readRand(jobChan, resultChan)
	}
	//循环打印数随机生成树的各位之和
	for value := range resultChan {
		fmt.Println(value)
	}
	wg.Wait()

}
