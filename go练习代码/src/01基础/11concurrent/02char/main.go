package main

import (
	"fmt"
	"time"
)

func CeShi() {
	var ch1 chan int
	ch1 = make(chan int, 10)
	ch1 <- 10
	ch1 <- 20
	ch1 <- 30
	x := <-ch1
	fmt.Println(x)
	close(ch1)
	<-ch1
	<-ch1
	<-ch1
	<-ch1
	<-ch1
	x1 := <-ch1
	fmt.Println(x1)
}

func recv(ch chan int) {
	x := <-ch
	fmt.Println(x)
}

func send(ch chan int) {
	ch <- 10
	fmt.Println("send success")
}

func new() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
	go send(ch)
	x := <-ch
	fmt.Println(x)

}

//比较两种循环
func forRange() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i2 := range ch2 {
		fmt.Println(i2)
	}
}

func selectIo() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case x := <-ch:
			fmt.Println(x)
		default:
			fmt.Println("没有任何操作")
		}

	}
}

func work(id int, jobs <-chan int, results chan<- int) {
	for i2 := range jobs {
		fmt.Printf("worker %d start job %d\n", id, i2)
		time.Sleep(time.Second)
		fmt.Printf("worker %d end job %d\n", id, i2)
		results <- i2 * 2
	}
}

func workPool() {
	jobs := make(chan int)
	results := make(chan int, 5)
	for i := 0; i < 3; i++ {
		go work(i, jobs, results)
	}
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
}

func main() {
	//new()
	//forRange()
	//selectIo()
	workPool()
}
