package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	id    int
	age   int8
	score int
	name  string
}

type Studentmanagement struct {
	students []*student
}

func NewStudent(id int, age int8, score int, name string) *student {
	return &student{
		id:    id,
		age:   age,
		name:  name,
		score: score,
	}
}

func (Sm *Studentmanagement) add(student *student) {
	Sm.students = append(Sm.students, student)
}

func (Sm *Studentmanagement) show() {
	fmt.Println("------------------")
	for _, stu := range Sm.students {
		fmt.Println(stu.name, stu.id, stu.score, stu.age)
	}
}
func (Sm *Studentmanagement) delete(name string) {
	var index = 0
	for i, stu := range Sm.students {
		if stu.name == name {
			index = i
		}
	}
	Sm.students = append(Sm.students[:index], Sm.students[index+1:]...)
}

func (Sm *Studentmanagement) editAge(name string, newage int8) {
	for _, stu := range Sm.students {
		if stu.name == name {
			stu.age = newage
		}
	}
}
func main() {
	Stu := Studentmanagement{
		make([]*student, 0, 100),
	}
	for i := 1; i < 10; i++ {
		student := NewStudent(i, 18, rand.Intn(100), fmt.Sprintf("stu%02d", i))
		fmt.Println(student)
		Stu.add(student)
	}
	Stu.show()
	Stu.delete("stu04")
	Stu.show()
	Stu.editAge("stu05", 19)
	Stu.show()
}
