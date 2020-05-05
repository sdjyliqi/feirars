package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

//ChannelsArgs ...ChannelsArgs统计体
type ChannelsArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
}

func HandleChannels(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs ChannelsArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.ModuleName == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "uninstall":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "active":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "news":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "preserve":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "feirar":
		items, err := PingbackCenter.GetActiveChannel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}
