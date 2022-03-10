package main

import (
	"fmt"
)

func stuctDefine() {
	type student struct {
		name  string
		age   int
		class string
	}

	type stu struct {
		name, class string
		age         int
	}
}

func instantiation() {
	type student struct {
		name  string
		age   int
		class string
	}
	//第一种普通的实例化
	var stu1 student
	stu1.age = 18
	stu1.name = "李四"
	stu1.class = "一年三班"
	fmt.Printf("stu1 type=%T,stu1=%#v\n", stu1, stu1)
	//第二种实例化 指针类型
	var stu2 = &student{}
	fmt.Printf("stu2 type=%T,stu2=%#v\n", stu2, stu2)
	//第三种实例化 指针类型
	var stu3 = new(student)
	stu3.name = "旺旺"
	fmt.Printf("stu3 type=%T,stu3=%#v\n", stu3, stu3)
}

func initStruct() {
	type student struct {
		name  string
		age   int
		class string
	}
	//第一种使用键值对初始化
	//对顺序没有要求
	var stu01 = student{
		name: "周星星",
		age:  18,
	}
	fmt.Printf("%#v \n", stu01)
	//第二种使用值列表,必须严格定义的顺序,不能少字段
	var stu2 = student{
		"张三",
		18,
		"三年二班",
	}
	fmt.Printf("%#v \n", stu2)
}

func zuoye1()  {
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Printf("%p \n",&stus[2])
	for _, stu := range stus {
		fmt.Println(stu)
		fmt.Printf("%p",&stu)
		m[stu.name] = &stu

	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
type student struct {
	name  string
	class string
	age   int

}

func newStudent(name,class string,age int) *student {
	return &student{
		name:name,
		class:class,
		age:age,
	}
}

func aa()  {
	p1 := newStudent("zhou","计科12-1",19)
	fmt.Printf("%#v \n",p1)
}


func main() {
	/*
	var ss struct{
		name string
		age int
	}
	fmt.Println(ss.name)
	initStruct()
	instantiation()
	zuoye1()
		aa()
	 */
	//zuoye1()


}
