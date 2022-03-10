package main

import "fmt"

type Dog struct {
	name string
}

func (D Dog) Say() {
	fmt.Printf("%s会汪汪汪的叫 \n", D.name)
}

type Cat struct {
	name string
}

func (c Cat) Say() {
	fmt.Printf("%s会喵喵喵的叫 \n", c.name)
}

type Animal interface {
	Say()
}

func MakeHungry(a Animal) {
	a.Say()
}

func main() {
	var a1 Animal
	d1 := Dog{
		name: "乐乐",
	}
	a1 = d1
	MakeHungry(a1)
	c1 := Cat{name: "九月"}
	a1 =c1
	MakeHungry(a1)
}
