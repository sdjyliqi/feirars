package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/feirars/utils"
	"net/http"
)

//LoginArgs ...登录的请求体
type LoginArgs struct {
	UserName string `json:"name" form:"name" binding:"required"`
	Passport string `json:"passport" form:"passport" binding:"required"`
}

func UCLogin(c *gin.Context) {
	cityCode := c.GetHeader("IPLOC")
	provinceCode := c.GetHeader("IPPROVINCE")
	var reqArgs LoginArgs
	err := c.ShouldBind(&reqArgs)
	err = UCenter.Login(reqArgs.UserName, reqArgs.Passport)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "402", "msg": err.Error(), "city": cityCode})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "token": utils.CreateToken(reqArgs.UserName, reqArgs.Passport), "city": cityCode, "province": provinceCode})
}
