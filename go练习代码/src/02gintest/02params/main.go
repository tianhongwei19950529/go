package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web/query", func(c *gin.Context) {
		//获取浏览器的参数
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})

	r.GET("/web/DefaultQuery", func(c *gin.Context) {
		//获取浏览器的参数
		name := c.DefaultQuery("name", "onebody")
		age := c.DefaultQuery("age", "0")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})

	r.POST("/web/DefaultQuery", func(c *gin.Context) {
		//获取浏览器的参数
		name := c.DefaultQuery("name", "onebody")
		age := c.DefaultQuery("age", "0")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})

	r.GET("/web/GetQuery", func(c *gin.Context) {
		//获取浏览器的参数
		name, ok := c.GetQuery("name")
		if !ok {
			c.JSON(200, gin.H{"code": 101, "data": "缺少name参数"})
			return
		}
		age, ok := c.GetQuery("age")
		if !ok {
			c.JSON(200, gin.H{"code": 101, "data": "缺少age参数"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})
	//获取form表单数据
	r.POST("/web/post", func(c *gin.Context) {
		name := c.PostForm("name")
		if name != "小王子" {
			c.JSON(200, gin.H{"code": 100, "data": "login fail", "msg": ""})
			return
		}
		c.JSON(200, gin.H{"code": 100, "data": "login successful", "msg": ""})
	})

	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{"name": name, "age": age})
	})

	r.GET("blog/:year/:mon", func(c *gin.Context) {
		year := c.Param("year")
		mon := c.Param("mon")
		c.JSON(200, gin.H{"year": year, "mon": mon})
	})

	r.GET("blog/:year/12", func(c *gin.Context) {
		year := c.Param("year")
		c.JSON(200, gin.H{"year": year})
	})

	//使用shouldBind方式获取全部信息
	type Stu struct {
		Username string   `json:"username" yaml:"username" form:"username"`
		Hobble   []string `json:"hobble" form:"hobble"`
	}

	r.POST("/json", func(c *gin.Context) {
		var a Stu
		err := c.ShouldBind(&a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "err": err.Error()})
		} else {
			fmt.Println(a.Hobble)
			for _, v := range a.Hobble {
				fmt.Println(v)
			}
			c.JSON(http.StatusOK, gin.H{"username": a.Username, "hobble": a.Hobble})
		}
	})

	r.GET("/index", func(c *gin.Context) {
		c.Redirect(301, "https://baidu.com")
	})

	//跳转
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
		fmt.Println(c.Request.Header.Get("Accept-Encoding"))
		return
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "data": "123"})
	})
	r.Run()
}
