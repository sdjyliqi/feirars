package handle

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
)

//PingbackArgs ...pingback统计体
type PingbackArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	PageID     int    `json:"page" form:"page" binding:"required"`
	PageCount  int    `json:"page" form:"pcount" binding:"required"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
	Channels   string `json:"chn" form:"chn"`
	Name       string `json:"name" form:"chn"`
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
	case "feirar-news":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, count, err := PingbackCenter.GetFeirarNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	case "feirar-update":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, count, err := PingbackCenter.GetFeirarUpdateDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}

//下载excel
func HandleExcel(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	//var reqArgs PingbackArgs
	//err := c.ShouldBind(&reqArgs)
	//if err != nil || reqArgs.PageID <= 0 || reqArgs.PageCount <= 0 || reqArgs.TimeEnd <= 0 {
	//	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
	//	return
	//}
	fileName := "./excelfiles/liqi9999.csv"
	//c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=liqi9999.xlsx")
	//c.Header("Content-Transfer-Encoding", "binary")
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	line := []string{"编号", "姓名", "性别"}
	r2.Write(line)
	for i := 0; i < 10; i++ {
		line := []string{fmt.Sprintf("%d", i), "liqi dyyyyy", "M"}
		r2.Write(line)

	}
	r2.Flush()
	fmt.Print(buf)
	fout, err := os.Create(fileName)
	fout.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	fmt.Println(err)
	fout.Write(buf.Bytes())
	defer fout.Close()
	c.File(fileName)
	return
	//
	//switch reqArgs.ModuleName {
	//case "install":
	//	cols := PingbackCenter.GetInstallDetailCols()
	//	items, count, err := PingbackCenter.GetInstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	//导出excel
	//	xlsx := excelize.NewFile()
	//
	//
	//	_ = xlsx.Write(c.Writer)
	//	return
	//case "uninstall":
	//	cols := PingbackCenter.GetUninstallDetailCols()
	//	items, count, err := PingbackCenter.GetUninstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//
	//case "active":
	//	cols := PingbackCenter.GetActiveDetailCols()
	//	items, count, err := PingbackCenter.GetActiveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//case "news":
	//	cols := PingbackCenter.GetNewsDetailCols()
	//	items, count, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//
	//case "preserve":
	//	cols := PingbackCenter.GetPreserveDetailCols()
	//	items, count, err := PingbackCenter.GetPreserveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//case "feirar":
	//	cols := PingbackCenter.GetFeirarDetailCols()
	//	items, count, err := PingbackCenter.GetFeirarDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//case "feirar-news":
	//	cols := PingbackCenter.GetFeirarDetailCols()
	//	items, count, err := PingbackCenter.GetFeirarNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//case "feirar-update":
	//	cols := PingbackCenter.GetFeirarDetailCols()
	//	items, count, err := PingbackCenter.GetFeirarUpdateDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items, "total": count})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}
