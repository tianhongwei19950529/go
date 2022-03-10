package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json1", func(c *gin.Context) {
		data1 := gin.H{"name": "小王子", "hobble": []string{"篮球", "羽毛球", "台球"}, "age": 18}
		c.JSON(200, data1)
	})
	type res struct {
		Name   string `json:"name"`
		Age    int
		Hobble []string
	}
	r.GET("/json2", func(c *gin.Context) {
		data := res{Name: "小孩纸", Age: 18, Hobble: []string{"篮球", "羽毛球", "台球"}}
		c.JSON(200, data)
	})
	r.Run()
}
