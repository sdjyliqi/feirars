package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func UCLogin(c *gin.Context) {
	name,passport := c.PostForm("name"),c.PostForm("passport")
    err := UserCenter.Login(name,passport)
    if err == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
}
