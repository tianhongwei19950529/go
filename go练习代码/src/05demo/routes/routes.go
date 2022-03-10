package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func SetUp() *gin.Engine{
	r := gin.New()

	r.Use(gin.Recovery())
	r.GET("/index", func (c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":viper.GetString("app.name")})
	})
	return r
}