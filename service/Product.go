package service

import (
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	result := srv.AddProduct(c)
	c.JSON(200, result)
}

func GetProduct(c *gin.Context) {
	result := srv.GetProduct(c)
	c.JSON(200, result)
}
