package main

import "fmt"

func swithDemo4(age int) {
	switch {
	case age >= 40:
		fmt.Println("社会毒打")
	case age >= 18:
		fmt.Println("好好工作")
	case age > 0:
		fmt.Println("好好学习")
	default:
		fmt.Println("输入有误")
	}
}

func gotoDemo1() {
	for i := 0; i < 10; i++ {
		for x := 0; x < 10; x++ {
			if x == 2 {
				break
			}
			fmt.Println(x)
		}
	}
	return
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func sum99() {
	//打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", i, j, i*j)
		}
		fmt.Println()
	}
}
func main() {
	//score := 85
	//if score >= 95 {
	//	fmt.Println("A+")
	//} else if score >= 80 {
	//	fmt.Println("A")
	//} else {
	//	fmt.Println("B")
	//}
	//fmt.Println()
	//if score1 := 100; score1 >= 95 {
	//	fmt.Println("A+")
	//} else if score1 >= 80 {
	//	fmt.Println("A")
	//} else {
	//	fmt.Println("B")
	//}
	//for a := 0; a < 10; a++ {
	//	fmt.Println(a)
	//}
	//
	//fmt.Println()
	//
	//i := 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println( )

	//i1 := 0
	//for i1 < 10{
	//	fmt.Println(i1)
	//	i1++
	//}
	//
	//for {
	//	fmt.Println("hello world")
	//}
	//nums := 5
	//switch nums {
	//case 1:fmt.Println("大拇指")
	//case 2:fmt.Println("食指")
	//case 3:fmt.Println("中指")
	//case 4:fmt.Println("无名指")
	//case 5:fmt.Println("小拇指")
	//default:
	//	fmt.Println("输入有误")
	//}

	//switch nums := 7; nums {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("奇数")
	//case 0, 2, 4, 8:
	//	fmt.Println("偶数")
	//}
	//age := 17
	//swithDemo4(age)
	//gotoDemo1()
	sum99()
}
