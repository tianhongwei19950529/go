package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学院管理系统,请选择你要做的操作")
	fmt.Println("1.展示学生")
	fmt.Println("2.添加学生")
	fmt.Println("3.修改学生")
	fmt.Println("4.删除学生")
	fmt.Println("5.退出")
}

func inputNum() (operaNum int) {
	fmt.Printf("请输入你要做的操作:")
	fmt.Scanf("%d\n", &operaNum)
	return

}
func inputStu() (stu *student) {
	var (
		id    int
		name  string
		class string
	)
	fmt.Printf("按照以下规范输入:")
	fmt.Printf("请输入你要操作的学生的编号:")
	fmt.Scanf("%d\n", &id)
	fmt.Printf("请输入你要操作的学生的姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入你要操作的学生的班级:")
	fmt.Scanf("%s\n", &class)
	stu = NewStudent(id, name, class)
	return
}

func main() {
	//1.展示菜单
	NSM := NewStudentManagement()
	for {
		showMenu()
		num := inputNum()
		switch num {
		case 1:
			fmt.Println("1.展示学生")
			NSM.showAll()
		case 2:
			fmt.Println("2.添加学生")
			stu := inputStu()
			NSM.addStu(stu)
		case 3:
			fmt.Println("3.修改学生")
			stu := inputStu()
			NSM.editStu(stu)
		case 4:
			fmt.Println("4.修改学生")
			var stuId int
			fmt.Printf("请输入你要操作的学生的编号:")
			fmt.Scanf("%d\n", &stuId)
			NSM.deleteStu(stuId)
		case 5:
			os.Exit(0)
		}
	}

}
