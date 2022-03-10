1.  可见性	

    1.  如果一个GO语言中定义的字段大写,别的包就可以找到,小写的话就找不到
    2.  除了当前文件,所有的包都算是外部包,所以类似json这种包进行序列化的时候也要大写字段,否则就会出现json序列化出不来的情况

2.  json序列化与反序列化

    1.  json反序列化的时候需要找一个额外的参数去接受值,并且返回只有err

    2.  打我打我

        ```go
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
        	Title    string `json:"title"`
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
        
        ```

    3.  结构体标签

        1.  可以设置多个标签,标签中间用空格分隔

        ```go
        type class struct {
        	Title    string `json:"title" yaml:"title"`
        	StudentS []*student
        }
        ```

        

