package main

import (
	"errors"
	"fmt"
)

// Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
// 异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理.

// 1、内置函数 panic
// 2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
// 3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
// 4、直到goroutine整个退出，并报告错误

// 1、内置函数 recover
// 2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
// 3、一般的调用建议
// a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
// b). 可以获取通过panic传递的error

// 注意：
// 1.利用recover处理panic指令，**defer 必须放在 panic 之前定义(可以放在不是panic的函数中)，另外 recover 只有在 defer 调用的函数中才有效。**否则当panic时，recover无法捕获到panic，无法防止panic扩散。
// 2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
// 3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
func test() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		println(err.(string)) // 将 interface 转型为具体类型
	// 	}
	// }()
	panic("panic error!")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		println(err.(string)) // 将 interface 转型为具体类型
	// 	}
	// }()
}

func main1() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface 转型为具体类型
		}
	}()
	test()
	fmt.Println("fuck") // 此行不会执行，会直接执行 defer
}

// 延迟调用中的 panic
func main2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fuck")
			fmt.Println(err.(string))
		}
		panic("defer panic")
	}()
	panic("test panic")
}

// main3
// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回nil. 任何未捕获的错误都会延调用堆栈向外传递
func main3() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	defer recover()              // 无效
	defer fmt.Println(recover()) // 无效
	defer func() {
		func() {
			println("defer inner")
			recover()
		}()
	}()
	panic("test panic")
}

// 使用延迟函数时是有效的
func main4() {
	defer expect()
	panic("test panic")
}

func expect() {
	fmt.Println(recover())
}

// 执行需要保护的代码段，需要将其冲构成匿名函数！！！来确保后面代码正常执行

func main5(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				fmt.Println("rec...")
				z = 0
			}
		}()
		// panic("test panic")
		z = x / y
		return
	}()
	fmt.Println(z)
}

// 除了用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数的调用状态
// 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型

var ErrDivByZero = errors.New("division by zero")

func main6(x, y int) {
	switch z, err := div(x, y); err { // NOTE: switch 判断 error !!!
	case nil:
		println(z)
	case ErrDivByZero:
		fmt.Println(err.Error())
		panic(err)
	}
}

func div(x int, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

// go 实现类似 try catch 的异常处理

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main7() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5(4, 2)
	// main5(4, 0)
	// main6(10, 1)
	// main6(10, 0)
	main7()
}
