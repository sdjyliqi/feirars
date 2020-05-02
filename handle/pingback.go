package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

type PingbackArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	PageID     int    `json:"page" form:"page" binding:"required"`
	PageCount  int    `json:"page" form:"pcount" binding:"required"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
}

func HandlePingbak(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var requestArgs PingbackArgs
	err := c.ShouldBind(&requestArgs)
	fmt.Println(err, requestArgs)
	//evType := c.DefaultQuery("type", "install")
	//strPageID := c.DefaultQuery("page", "1")
	//strPageCount := c.DefaultQuery("pcount", "20")
	//pageID, errPageID := strconv.Atoi(strPageID)
	//pageCount, errpageCount := strconv.Atoi(strPageCount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch requestArgs.ModuleName {
	case "install":
		cols := PingbackCenter.GetInstallDetailCols()
		items, count, err := PingbackCenter.GetInstallDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "uninstall":
		cols := PingbackCenter.GetUninstallDetailCols()
		items, count, err := PingbackCenter.GetUninstallDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return

	case "active":
		cols := PingbackCenter.GetActiveDetailCols()
		items, count, err := PingbackCenter.GetActiveDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return

	case "news":
		cols := PingbackCenter.GetNewsDetailCols()
		items, count, err := PingbackCenter.GetNewsDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return

	case "preserve":
		cols := PingbackCenter.GetPreserveDetailCols()
		items, count, err := PingbackCenter.GetPreserveDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "feirar":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, count, err := PingbackCenter.GetFeirarDetailItems(requestArgs.PageID, requestArgs.PageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}

func HandlePingbakInstall(c *gin.Context, pageID, pageCount int) {
	cols := PingbackCenter.GetInstallDetailCols()
	items, count, err := PingbackCenter.GetInstallDetailItems(pageID, pageCount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
}
