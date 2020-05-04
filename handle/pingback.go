package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

//PingbackArgs ...pingback统计体
type PingbackArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	PageID     int    `json:"page" form:"page" binding:"required"`
	PageCount  int    `json:"page" form:"pcount" binding:"required"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
	Channels   string `json:"chn" form:"chn"`
}

func HandlePingbak(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs PingbackArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.PageID <= 0 || reqArgs.PageCount <= 0 || reqArgs.TimeEnd <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		cols := PingbackCenter.GetInstallDetailCols()
		items, count, err := PingbackCenter.GetInstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "uninstall":
		cols := PingbackCenter.GetUninstallDetailCols()
		items, count, err := PingbackCenter.GetUninstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return

	case "active":
		cols := PingbackCenter.GetActiveDetailCols()
		items, count, err := PingbackCenter.GetActiveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "news":
		cols := PingbackCenter.GetNewsDetailCols()
		items, count, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return

	case "preserve":
		cols := PingbackCenter.GetPreserveDetailCols()
		items, count, err := PingbackCenter.GetPreserveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "feirar":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, count, err := PingbackCenter.GetFeirarDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}
