package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//func createRouter(useDefault bool) (r *gin.Engine) {
//	if useDefault {
//		r = gin.Default()
//	} else {
//		r = gin.New()
//		r.Use(gin.Logger())
//		r.Use(gin.Recovery())
//	}
//	return r
//}

func getRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)
	v2 := router.Group("/v2")
	addPingRoutes(v2)
}

func Run() {
	var (
		router = gin.Default()
	)
	getRoutes(router)
	err := router.Run(":5021")
	if err != nil {
		fmt.Println(err)
		return
	}
}
