package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//ChartArgs ... 曲线区统计体
type ChartArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
	Channels   string `json:"chn" form:"chn"`
}

func HandleChart(c *gin.Context) {
	var reqArgs ChartArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.ModuleName == "" || reqArgs.TimeStart < 0 || reqArgs.TimeEnd < 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误1111"})
		return
	}
	switch reqArgs.ModuleName {

	case "feirar":
		items, err := PingbackCenter.GetFeirarChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}
