package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/feirars/handle"
	"github.com/sdjyliqi/feirars/middleware"
)

func InitRouter(r *gin.Engine) {

	r.Use(middleware.Cors())
	// uc先关接口
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/pingback", handle.HandlePingbak)
		GroupV1.GET("/login", handle.UCLogin)
	}

}
