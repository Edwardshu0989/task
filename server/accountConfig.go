package server

import (
	"github.com/gin-gonic/gin"
)

func AddAccountConfig(c *gin.Context) (err error) {
	c.GetPostForm("action")
	err = nil
	return
}

//func RsaCommodityInfo(c *gin.Context) {
//	commodityInfo := c.GetPostForm("commodity_info")
//	rsa := sha256.
//}
