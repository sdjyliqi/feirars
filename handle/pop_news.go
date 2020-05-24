package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

type PopNewsArgs struct {
	UID       string `json:"uid" form:"uid"` //binding:"required"
	UNX       string `json:"unx" form:"unx"`
	Ver       string `json:"ver" form:"ver" `
	Frm       string `json:"frm" form:"frm" `
	SoftID    int64  `json:"softid" form:"softid"`
	OS        string `json:"os" form:"os"`
	Start     string `json:"start" form:"start"`
	EventType string `json:"type" form:"type"`
	City      string `json:"city" form:"city"`
}

func HandlePopNews(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	fmt.Println(header)
	var reqArgs PopNewsArgs
	err := c.ShouldBind(&reqArgs)
	reqArgs.City = c.GetHeader("IPCITY")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": reqArgs.City})
		return
	}
	//利用header的值，确定city的值。
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": reqArgs.City})
}
