package main

// golang方法集：每个类型都有与之关联的方法集，这会影响到接口的实现规则
// • 类型 T 方法集包含全部 receiver T 方法。
// • 类型 *T 方法集包含全部 receiver T + *T 方法。
// • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
// • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
// • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。

// 用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，编译器总是查找全部方法，并自动转换 receiver 实参。
import (
	"fmt"
)

type T struct {
	int
}

type S struct {
	T
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func main() {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.test()
	t1.testP()

	t2 := &t1
	t2.test()
	t2.testP()

	s1 := S{T{1}}
	fmt.Printf("s1 is %v\n", s1)
	s1.test()
	s1.testP()

	s2 := &s1
	fmt.Printf("s2 is %v\n", s2)
	s2.test()
	s2.testP()
}
