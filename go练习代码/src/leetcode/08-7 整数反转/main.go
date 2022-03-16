package main

import (
	"fmt"
)

/*
我想到的是字符串反转
但是可以用数学的方式来解决,求模取余  反向相加
*/

func reverse(x int) (rev int) {
	for x != 0 {
		//if rev < (1<<31) || rev > (1>>31) {
		//	fmt.Println("------------")
		//	return 0
		//}
		//if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
		//	return 0
		//}
		dis := x % 10
		x = x / 10
		rev = rev*10 + dis
	}
	fmt.Println(rev)
	return rev
}

func main() {
	a := -138
	fmt.Println(reverse(a))
}
