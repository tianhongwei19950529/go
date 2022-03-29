package main

import (
	"context"
	"fmt"
<<<<<<< HEAD
	"strconv"
=======
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f
	"sync"
	"time"
)

// 利用 sync.Group 实现协程同步
<<<<<<< HEAD
// WaitGroup内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：
// Add()用来添加计数。
// Done()用来在操作结束时调用，使计数减一。
// Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回。
=======
//WaitGroup内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：
//Add()用来添加计数。
//Done()用来在操作结束时调用，使计数减一。
//Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回。
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f

var wg sync.WaitGroup
var exit bool

// demo1: 使用 waitGroup 进行计数
func worker1() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func main1() {
	const Num = 3
	for i := 0; i < Num; i++ {
		wg.Add(1)
		go worker1()
	}
	time.Sleep(time.Second * 3)
	exit = true
	wg.Wait()
}

// demo2: 使用 channel
func main2() {
	const Num = 3
	stopCh := make(chan struct{})
	for i := 0; i < Num; i++ {
		go worker2(stopCh)
	}
	time.Sleep(time.Second * 3)
	stopCh <- struct{}{}
	fmt.Println("over...")
	// 如何保证goroutine全部退出？
}

func worker2(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
			fmt.Println("worker2")
			time.Sleep(time.Second)
		}
	}
}

// demo3: 用 context 代替 channel
// context.Background() 返回一个空 Context，它一般用于整个 Context 树的根节点。
// context.WithCancel() 创建一个可取消的子 Context，返回 ctx 和 cancel
// <-ctx.Done() 判断是否要结束，在调用 cancel 之后，会一直收到信号
// ctx.Err()  返回取消的错误原因
// context.Value() 获取该 Context 上绑定的值，是一个键值对，这个值是线程安全
// Deadline 设置截止时间，到该时间点便会自动 cancel
// Timeout 超时自动取消

<<<<<<< HEAD
// func worker3(ctx context.Context) {
=======
//func worker3(ctx context.Context) {
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f
//	for {
//		select {
//		case <-ctx.Done(): // cancel 调用后，这个case会疯狂执行，一直能收到数据
//			fmt.Println("err", ctx.Err()) // 取消的原因
//			fmt.Println("v", ctx.Value("key"))
//			fmt.Println("cancel...")
//			wg.Done()
//			return
//		default:
//			fmt.Println("worker")
//			time.Sleep(time.Second * 3)
//		}
//	}
<<<<<<< HEAD
// }
=======
//}
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f

func worker3(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done(): // cancel 调用后，这个case会疯狂执行，一直能收到数据
			fmt.Println("err", ctx.Err()) // 取消的原因
			fmt.Println("v", ctx.Value("key"))
			fmt.Println("cancel...")
			break LOOP // 艹特么的 beak label
		default:
			fmt.Println("worker")
			time.Sleep(time.Second * 3)
		}
	}
	wg.Done()
}

func main3() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "key", "监控1")
	wg.Add(1) // 尽量在外面加1，防止 goroutine 没启动直接跑 wait
	go worker3(ctx2)
	time.Sleep(time.Second * 3)
	cancel() // 通知子 goroutine 结束 Done
	wg.Wait()
	fmt.Println("over..")
}

// WithDeadline
func main4() {
	d := time.Now().Add(500 * time.Millisecond) // time.Now().Add()是什么鬼？当前时间多久之后嘛
<<<<<<< HEAD
	// ctx, _ := context.WithDeadline(context.Background(), d)
=======
	//ctx, _ := context.WithDeadline(context.Background(), d)
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管 ctx 会过期，但在任何情况下调用它的 cancel 函数都是很好的实践
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	cancel()
	cancel()
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// WithTimeout
func main5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	wg.Add(1)
	go worker3(ctx)
<<<<<<< HEAD
	// cancel()
=======
	//cancel()
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f
	time.Sleep(10 * time.Second)
	wg.Wait()
	fmt.Println("over")
}

<<<<<<< HEAD
// 互斥锁
var x int64
var swg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock() // 不加锁，每次打出的x都是不一样的
		x = x + 1
		lock.Unlock()
	}
	swg.Done()
}

func main6() {
	num := 1000
	swg.Add(num)
	for i := 0; i < num; i++ {
		go add()
	}
	swg.Wait()
	fmt.Println(x)
}

// 读写互斥锁 sync.RWMutex

func main7() {
	var wwg sync.WaitGroup
	var rwlock sync.RWMutex
	go func() { // 默认写
		rwlock.Lock()
		time.Sleep(time.Second)
		rwlock.Unlock()
		wwg.Done()
	}()
	wwg.Add(1)
	go func() { // 读
		rwlock.RLock()
		time.Sleep(time.Second)
		rwlock.RUnlock()
		wwg.Done()
	}()
	wwg.Add(1)
	wwg.Wait()
}

// sync.Once 高并发场景下只执行一次，例如只加载一次的配置文件，只关闭一次的通道等
// var icons map[string]image.Image
//
// var loadIconsOnce sync.Once
//
// func loadIcons() {
// 	icons = map[string]image.Image{
// 		"left":  loadIcon("left.png"),
// 		"up":    loadIcon("up.png"),
// 		"right": loadIcon("right.png"),
// 		"down":  loadIcon("down.png"),
// 	}
// }
//
// // Icon 是并发安全的
// func Icon(name string) image.Image {
// 	loadIconsOnce.Do(loadIcons)
// 	return icons[name]
// }

// sync.Map 线程安全的map

func main8() {
	wg := sync.WaitGroup{}
	var m sync.Map
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m.Store(key, i)
			value, _ := m.Load(key)
			fmt.Printf("k:%v, v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	main8()
=======
func main() {
	//main1()
	//main2()
	//main3()
	//main4()
	main5()
>>>>>>> e1ba3013c3a2ca15e1a71d8865d02cdb2404e39f
}
