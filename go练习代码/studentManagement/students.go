package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

type studentManagment struct {
	allstudent []*student
}

// NewStudent 学生的构造函数
func NewStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// NewStudentManagement 系统的构造函数
func NewStudentManagement() *studentManagment {
	return &studentManagment{
		allstudent: make([]*student, 0, 20),
	}
}

//添加学生
func (sM *studentManagment) addStu(stu *student) {
	sM.allstudent = append(sM.allstudent, stu)
}

//展示学生
func (sM *studentManagment) showAll() {
	for _, stu := range sM.allstudent {
		fmt.Println(stu.id, stu.name, stu.class)
	}
}

func (sM *studentManagment) editStu(stu *student) {
	for i, stu1 := range sM.allstudent {
		if stu1.id == stu.id {
			sM.allstudent[i] = stu
		}
	}
}

func (sM *studentManagment) deleteStu(xuehao int) {
	i1 := -1
	for i, stu1 := range sM.allstudent {
		fmt.Println(stu1.id==11)
		if stu1.id == xuehao {
			i1 = i
		}
	}
	if i1 == -1 {
		fmt.Println("没有这个人")
		return
	}
	fmt.Println(i1)
	sM.allstudent = append(sM.allstudent[:i1], sM.allstudent[i1+1:]...)
}
