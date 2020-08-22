package resultdata

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandNotFond(c *gin.Context) {
	c.JSON(http.StatusOK, "Router Err")
}
