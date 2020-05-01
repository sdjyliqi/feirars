package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/satori/go.uuid"
)

//... RequestIDMiddleware  add request id into header
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		glog.Info("add id in header done:", c.Request.Header)
		c.Next()
		glog.Infof("response header:%+v,code:%+v,len:", c.Writer.Header(), c.Writer.Status(), c.Writer.Size())
	}
}
