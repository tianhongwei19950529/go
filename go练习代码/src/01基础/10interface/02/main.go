package main

import "fmt"

type stu struct {
	name string
	age  int
}

type tea struct {
	name string
	age  int
}

func (s *stu) addAge(age int) {
	s.age = s.age + 10
}

func (t *tea) addAge(age int) {
	t.age = t.age + 10
}

type age interface {
	addAge(int)
}

func newAge(age age, num int) {
	age.addAge(num)
}

func main() {
	s := stu{
		name: "wang",
		age:  10,
	}
	newAge(&s, 20)
	fmt.Println(s.age)
}
