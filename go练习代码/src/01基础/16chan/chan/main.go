// channel的应用

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 一、停止信号（优雅关闭channel）
// 有一条广泛流传的关闭channel原则：不要从一个receiver侧关闭channel，也不要在有多个sender存在时，关闭channel
// 上面所说的不是最本质的，最本质的原则只有一条：不要关闭已经关闭的channel，或者向其发送数据
// 有两个不那么优雅的关闭channel的方法
// 1. 使用 defer-recover机制，放心大胆的关闭channel或者向channel发送数据。即使发生了panic，有defer-recover在兜底。
// 2. 使用 sync.Once 来保证只关闭一次
// 到底该如何优雅的关闭 channel? 根据 sender 和 receiver的个数，分下面几种情况：
// 1. 一个 sender，一个 receiver
// 2. 一个 sender，M 个 receiver
// 3. N 个 sender，一个 receiver
// 4. N 个 sender，M 个 receiver
// 对于 1, 2，只有一个 sender 情况就不用多说了，直接从 sender 端关闭就好了，重点关注 3, 4 中情况
// 第 3 种情形下，优雅关闭 channel 的方法是：增加一个传递关闭信号的 channel，receiver 通过信号下达关闭数据 channel 指令，sender 监听到关闭信号后，停止接收数据。
func main1_1() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return // 这个 return 写的很有灵魂！！！
				case dataCh <- rand.Intn(Max): // 不停地 发送值到通道
				}
			}
		}()
	}
	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signale to senders.")
				fmt.Println("vvv", value)
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()

	// 不写 for，执行完 case 就退出？time.After 返回个 chan Time
	select {
	case <-time.After(time.Hour):

	}

	// 对于 dataCh，上面的代码并没有关闭它！！！因为 goroutine 都已经执行完成，没有 goroutine 在引用它，所以不管它有没有被关闭，都会被 gc 回收。所以这时候优雅的关闭 chanel 就是不关闭，让 gc 代劳
}

// 第 4 种情况，有 M 个 receiver，N个 sender，需要增加一个中间人，M 个 receiver 都向他发送关闭 dataCh 的请求，中间人收到第一个请求后，就会下达关闭 dataCh 的指令（通过关闭stopCh，这时就不会发生重复关闭的情况，应为 stopCh 的发送方只有中间人一个），另外，N 个 sender 也可以向中间人发送关闭 dataCh 的请求。

func main1_2() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// It must be a buffered channel
	toStop := make(chan string, 1)
	var stoppedBy string

	go func() {
		stoppedBy = <-toStop
		fmt.Println(stoppedBy)
		close(stopCh)
	}()
	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select { // 为什么要用 select 加 default ? 因为可能多sender同时 value ==0 但进行阻塞。如果发生阻塞那么谁先产生的就不知道？
					case toStop <- "sender:" + id:
					default:
					}
					return
				}
				select {
				case <-stopCh:
					return
				case dataCh <- value:
					fmt.Println("sender#" + id + "pushed" + strconv.Itoa(value))
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					fmt.Println("receiver#" + id + "received" + strconv.Itoa(value))
					if value == Max-1 {
						select {
						case toStop <- "receiver:" + id:
						default:
						}
						return
					}
				}
			}
		}(strconv.Itoa(i))
	}
	fmt.Println("stop...")
	select {
	case <-time.After(time.Hour):
	}
}

// 加入 toStop 是个非缓冲性 channel，那么第一个发送的关闭 dataCh 请求可能丢失，因为无论是 sender 还是 receiver 都是通过 select 语句来发送请求，如果中间人所在的 goroutine 没准备好，那么 select 语句就不会选中，直接走 default 选项，什么也不做。这样第一个关闭 dataCh 的请求就会丢失。

// 二、任务定时
// 需要执行某项操作，但又不想它耗费太长时间
//func main2_1() {
//	select {
//	case <- time.After(100 * time.Millisecond):
//	case <- s.stopc:
//		return false
//	}
//}
// 定时执行某个任务
func main2_2() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行1s定时任务")
		}
	}
}

// 三、解耦生产方和消费方
func main3_1() {
	taskCh := make(chan int, 100)
	var taskNum int = 1000
	go worker(taskCh)
	for i := 0; i < taskNum; i++ {
		taskCh <- i
	}
	select {
	case <-time.After(time.Hour):
	}
}

func worker(taskCh <-chan int) {
	const N = 50
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

// 控制并发数
func main() {
	// ......
	var limit = make(chan int, 3)
	for i := 0; i < 1000; i++ {
		go func() {
			limit <- 1
			sayHello()
			<-limit
		}()
	}
}

func sayHello() {
	fmt.Println("hello")
	time.Sleep(time.Second)
}
