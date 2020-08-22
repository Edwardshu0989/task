package service

import (
	"awesomeProject/server"
	"github.com/gin-gonic/gin"
)

func AccountConfigInfoAdd(c *gin.Context) {
	err := server.AddAccountConfig(c)
	c.JSON(200, err)
}
