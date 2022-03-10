package main

import (
	"fmt"
	"strings"
)

func function1() {
	fmt.Println("hello")
}

func function2(a int) {
	fmt.Println(a)
}

func function3(a, b int) {
	fmt.Println(a + b)
}

func function4(a, b int) int {
	return a + b
}
func function5(a, b int) (res, difference int) {
	return a + b, a - b
}

func function6(a int, b ...[]int) (res int) {
	res = a
	for _, v := range b {
		for _, v1 := range v {
			fmt.Println(v1)
			res += v1
			v1 = v1 + 1
			fmt.Println(res)
			fmt.Println(v1)
		}

	}
	return
}

func add(x, y int) int {
	return x + y
}
func sub(m, n int) (aa int) {
	return m - n
}

func cala(x, y int, name func(m, n int) int) int {
	return name(x, y)
}

func do(s string) func(x, y int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// return语句执行步骤
// 1、返回值赋值
// 2、defer语句
// 3、真正RET返回
func f0() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return x //返回值RET=x, x++, RET=x=6
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //返回值RET=5, x++, RET=5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值RET=x=5, x++, RET=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值RET=y=x=5, x++, RET=5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值RET=x=5, x`++, RET=5
}

func bibao(suff string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suff) {
			return name
		}
		return name + suff
	}
}

func A() {
	fmt.Println("func A")
}

func B() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error in B")
		}
	}()
	//fmt.Println("func B")
	panic("func B")
}
func C() {
	fmt.Println("func C")
}

func dispatchCoin(coins int, users []string, distribution map[string]int) int {
	for _, names := range users {
		var bs = []rune(names)
		count := 0
		for _, name := range bs {
			if name == 'e' || name == 'E' {
				count += 1
			}else if  name == 'i' || name == 'I'{
				count +=2
			} else if  name == 'o' || name == 'O'{
				count +=3
			} else if  name == 'u' || name == 'U'{
				count +=4
			}
		}
		distribution[names] = count
		coins -= count
	}
	fmt.Println(distribution)
	return coins
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin(coins, users, distribution)
	fmt.Println("剩下：", left)
	//A()
	//B()
	//C()
	//r := bibao(".thw")
	//fmt.Println(r("aa"))

	//x := 1
	//y := 2
	//defer calc("AA", x, calc("A", x, y))
	//x = 10
	//defer calc("BB", x, calc("B", x, y))
	//y = 20
	//name := do("-")
	//res := name(10,20)
	//fmt.Println(res)
	//res := cala(10, 20, add)
	//fmt.Println(res)
	//type calculation func(x, y int) int
	//var a11 calculation
	//a11 = add
	//a11(10, 20)
	//a11 = sub
	//ret := func(){
	//	fmt.Println("hello")
	//}
	//ret()
}
