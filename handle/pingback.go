package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"strconv"
)

func HandlePingbak(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	evType := c.DefaultQuery("type", "install")
	strPageID := c.DefaultQuery("page", "1")
	strPageCount := c.DefaultQuery("pcount", "20")
	pageID, errPageID := strconv.Atoi(strPageID)
	pageCount, errpageCount := strconv.Atoi(strPageCount)
	if errPageID != nil || errpageCount != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch evType {
	case "install":
		cols := PingbackCenter.GetInstallDetailCols()
		items, count, err := PingbackCenter.GetInstallDetailItems(pageID, pageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "uninstall":
		cols := PingbackCenter.GetUninstallDetailCols()
		items, count, err := PingbackCenter.GetUninstallDetailItems(pageID, pageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})

	case "active":
		cols := PingbackCenter.GetActiveDetailCols()
		items, count, err := PingbackCenter.GetActiveDetailItems(pageID, pageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})

	case "news":
		cols := PingbackCenter.GetNewsDetailCols()
		items, count, err := PingbackCenter.GetNewsDetailItems(pageID, pageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})

	case "preserve":
		cols := PingbackCenter.GetPreserveDetailCols()
		items, count, err := PingbackCenter.GetPreserveDetailItems(pageID, pageCount)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	}

}

func HandlePingbakInstall(c *gin.Context, pageID, pageCount int) {
	cols := PingbackCenter.GetInstallDetailCols()
	items, count, err := PingbackCenter.GetInstallDetailItems(pageID, pageCount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
}
