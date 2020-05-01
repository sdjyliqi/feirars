package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HandleHi(c *gin.Context){
	header := c.Request.Header
	c.JSON(http.StatusOK,gin.H{"name": "hi....","time":time.Now()})
}
