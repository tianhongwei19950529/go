package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
	})

	shopGroup := r.Group("/shop")
	{
		shopGroup.Any("index", func(c *gin.Context) {
			switch c.Request.Method {
			case http.MethodGet:
				c.JSON(http.StatusOK, gin.H{"data": "/shop/index"})
			}
		})

		shopGroup.Any("woman", func(c *gin.Context) {
			switch c.Request.Method {
			case http.MethodGet:
				c.JSON(http.StatusOK, gin.H{"data": "/shop/woman"})
			}
		})

	}
	r.Run()
}
