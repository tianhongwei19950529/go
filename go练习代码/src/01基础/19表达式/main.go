package main

import "fmt"

// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
//
//    instance.method(args...) ---> <type>.func(instance, args...)
// 前者称为 method value，后者 method expression。
//
// 两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。

type User struct {
	id   int
	name string
}

func (u *User) TestPointer(fuck int) {
	fmt.Println("----", u)
	fmt.Println(fuck)
	fmt.Printf("%p, %v\n", u, u)
}

func (u User) TestValue(fuck int) {
	fmt.Println("----", u)
	fmt.Println(fuck)
	fmt.Printf("%p, %v\n", &u, u) // 内存地址全tmd不同！！！
}

func main() {
	u := User{1, "Tom"}

	// 隐式传递
	f := u.TestValue // 立即复制 receiver，因为不是指针类型，不受后续修改影响。
	u.id = 2
	f(10)

	// 显示传递
	mExpression := (*User).TestValue
	mExpression(&u, 10) // 传递我自己！！！
}
