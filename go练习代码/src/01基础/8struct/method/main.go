package main

import "fmt"

type student struct {
	name string
	age  int
}

func NewStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func (s student) Dream() {
	fmt.Printf("%s想要学号go语言！！", s.name)
}

func (s student) SetAge(newAge int) {
	s.age = newAge
}

func (s *student) SetAge2(newAge int) {
	s.age = newAge
}

func demo1() {
	s1 := NewStudent("孙悟空", 18)
	s1.Dream()
}

func demo2() {
	s1 := NewStudent("孙悟空", 18)
	fmt.Println("age:", s1.age)
	s1.SetAge(20)
	fmt.Println("age:", s1.age)
	s1.SetAge2(20)
	fmt.Println("age:", s1.age)
}

func main() {
	//demo1()
	demo2()
}
