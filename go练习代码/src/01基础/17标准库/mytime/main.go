package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 时间类型
func main1() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳
func main2() {
	now := time.Now()
	timestamp1 := now.Unix()     // int64 时间戳
	timestamp2 := now.UnixNano() // int64 时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 时间戳转时间格式
func tsftime(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // NOTE 将时间戳转化为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间间隔
func main3() {
	now := time.Now()                            // 当前时间 Time 对象
	later := now.Add(time.Hour).Add(time.Minute) // 运算之后的 Time 对象
	fmt.Println(now.AddDate(1, 1, 0))
	sub := now.Sub(later)
	dur := later.Sub(now) // 做减法的时间间隔 Duration，和 Add 之间的类型一样(int64别名)
	fmt.Println(sub)
	fmt.Println(dur)
	var l bool = later.After(now)
	var z bool = later.Before(now)
	fmt.Println(l, z)
	fmt.Println(later.Equal(later))
}

// 周期定时器 tick，类似 beat
func main4(wg *sync.WaitGroup) {
	ticker := time.Tick(time.Second)
	t := 0
	for i := range ticker {
		t++
		fmt.Println(i)
		if t > 10 {
			break
		}
	}
	wg.Done()
}

// 时间格式化
func strftime() {
	now := time.Now()
	fmt.Println(now.Format("2006-02-01 04:03:15"))
}

// 解析字符串格式的时间
func strptime() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-12-31 23:59:59", loc)
	fmt.Println(timeObj, err)
}

// 定时器 timer，只能响应一次；可停止、重置等
func main5() {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	// go func() {
	go func() {
		t2 := <-timer1.C // C means channel
		fmt.Printf("t2:%v\n", t2)
	}()

	// timer1 可以被停止，必须在定时器走完之前停止
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		b := timer1.Stop() // 定时器走完返回false，否则停止并返回true
		fmt.Println(b)
		if b {
			fmt.Println("time1 has been stop")
		} else {
			timer1.Reset(1 * time.Second)
			fmt.Println("reset now..", time.Now())
			fmt.Println(<-timer1.C)
		}
	}()

	// 还可以用 time.After

	<-time.After(3 * time.Second)
	t3 := time.Now()
	fmt.Printf("t2:%v\n", t3)

}

func main() {
	// main1()
	// main2()
	// tsftime(time.Now().Unix())
	// main3()
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go main4(&wg)
	// strftime()
	// wg.Wait()
	strptime()
	main5()
}
