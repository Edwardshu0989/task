package service

import (
	"awesomeProject/resultdata"
	v1 "awesomeProject/server"
	"github.com/gin-gonic/gin"
)

var (
	srv *v1.Server
)

func New(v *v1.Server) (engine *gin.Engine, err error) {
	// 设置为release mode
	gin.SetMode(gin.DebugMode)

	// 创建engine
	engine = gin.New()
	// 初始化路由
	initRouter(engine)
	srv = v
	err = engine.Run(":8081")
	return
}

func initRouter(v *gin.Engine) {
	v.NoMethod(resultdata.HandNotFond)
	v.NoRoute(resultdata.HandNotFond)
	v.GET("/ping", Ping)
	v1 := v.Group("/v1")
	{
		v1.POST("/add/redis", InfoAdd)
		v1.POST("/get/redis", InfoGet)
		v1.POST("/add/product", AddProduct)
		v1.POST("/get/product", GetProduct)
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, "hehe")
}
