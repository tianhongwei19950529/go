package main

import (
	"fmt"
)

func anonymous() {
	//测试匿名结构体
	type persion struct {
		string
		int
	}
	p1 := persion{
		"小王子",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.int, p1.string)
}

func nest1() {
	type address struct {
		provice string
		city    string
	}
	type persion struct {
		name string
		age  int
		 *address
	}
	p1 := persion{
		"小王子",
		18,
		&address{
			provice: "山东",
			city:    "济南",
		},
	}
	fmt.Printf("%#v \n", p1)
	fmt.Println(p1.address.provice)
	fmt.Println(p1.provice)
}

func nest2()  {
	//双重嵌套重复字段
	type address struct {
		provice string
		city    string
	}
	type email struct {
		provice string
		city    string
	}

	type persion struct {
		name string
		age  int
		*address
		email
	}

	p1 := persion{
		name:"小王子",
		age: 18,
		address:&address{
			provice: "山东",
			city:    "济南",
		},
		email:email{
			provice: "北京",
			city:"北京",
		},
	}
	fmt.Printf("%#v \n", p1)
	fmt.Println(p1.address.provice)
	fmt.Println(p1.email.provice)
}

type Animal struct {
	name string
}


type dog struct {
	*Animal
	//name string
	lastName string
}

func (a *Animal)move()  {
	fmt.Println(a.name,"会移动")
}

func (a *Animal)run()  {
	fmt.Println(a.name,"会跑")
}

func (d dog)wang()  {
	fmt.Println(d.name,"会叫")
}

func (d dog)move()  {
	fmt.Println(d.lastName,"会pao")
}
func nest3()  {
	d1 := dog{
		lastName: "乐乐",
		Animal:&Animal{
			name:"狗",
		},
	}
	d1.move()
	d1.wang()
	d1.run()
}
func main() {
	//anonymous()
	//nest1()
	//nest2()
	nest3()
}

