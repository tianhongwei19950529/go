package main

import (
	"fmt"
)

func main() {
	type xxx interface{}
	var a xxx
	a = 10
	fmt.Println(a)
	a = "hahaha"
	fmt.Println(a)
	a = make([]int, 0, 10)
	a = "牛逼"
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("....")
	}
	v, ok := a.(string)
	if ok {
		fmt.Println("你猜对了", v)
	} else {
		fmt.Println("你猜错了")
	}
}
