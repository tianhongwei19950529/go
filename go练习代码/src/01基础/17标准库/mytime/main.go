package main

import (
	"fmt"
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
	sub := now.Sub(later)
	dur := later.Sub(now) // 做减法的时间间隔 Duration，和 Add 之间的类型一样(int64别名)
	fmt.Println(sub)
	fmt.Println(dur)
	var l bool = later.After(now)
	var z bool = later.Before(now)
	fmt.Println(l, z)
	fmt.Println(later.Equal(later))
}

// 定时器
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

func main() {
	//main1()
	//main2()
	//tsftime(time.Now().Unix())
	//main3()
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go main4(&wg)
	//strftime()
	//wg.Wait()
	strptime()
}
