package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string `json:"title" yaml:"title"`
	StudentS []*student
}

func newStudent(id int, name string) *student {
	return &student{
		Name: name,
		ID:   id,
	}
}

func main() {
	c1 := class{
		Title:    "三年二班",
		StudentS: make([]*student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.StudentS = append(c1.StudentS, tmpStu)
	}

	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("err :", err)
	}
	fmt.Printf("%s \n", data)
	fmt.Println("-----------------------")

	json_str := `{"title":"三年二班","StudentS":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"}]}`
	var c2 class
	err = json.Unmarshal([]byte(json_str), &c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", c2)
}
