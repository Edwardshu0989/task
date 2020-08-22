package service

import (
	"awesomeProject/resultdata"
	"github.com/gin-gonic/gin"
)

func New() (engine *gin.Engine, err error) {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	initRouter(engine)
	err = engine.Run(":8080")
	return
}

func initRouter(v *gin.Engine) {
	v.NoMethod(resultdata.HandNotFond)
	v.NoRoute(resultdata.HandNotFond)
	v.GET("/ping", Ping)
	v1 := v.Group("/v1")
	{
		v1.POST("/add/config", AccountConfigInfoAdd)
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, "hehe")
}
