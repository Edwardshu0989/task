package service

import (
	"github.com/gin-gonic/gin"
)

func InfoAdd(c *gin.Context) {
	resp := srv.AddRedis()
	c.JSON(200, resp)
}

func InfoGet(c *gin.Context) {
	resp := srv.GetRedisData()
	c.JSON(200, resp)
}
