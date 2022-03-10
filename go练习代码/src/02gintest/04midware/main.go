package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func name(c *gin.Context) {
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{"msg": name})
}

func cost(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now().Sub(start)
	fmt.Println(end)
}

func m1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("m1")
		c.Set("name", "小王子")
		c.Next()
		fmt.Println("m2")
	}
}
func main() {
	r := gin.Default()
	r.Use(cost, m1())
	r.GET("/index", name)
	r.Run()
}
