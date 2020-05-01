package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"time"
)

func HandleHi(c *gin.Context){
	header := c.Request.Header
	glog.Info("Header in request:",header)
	c.JSON(http.StatusOK,gin.H{"name": "hi....","time":time.Now()})
}
