package main

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdjyliqi/feirars/middleware"
	"github.com/sdjyliqi/feirars/router"
)

func main() {
	r := gin.Default()
	r.Use(middleware.RequestIDMiddleware())
	r.Use(ginprom.PromMiddleware(nil))
	// register the `/metrics` route.
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	router.InitRouter(r)
	r.Run(":8844")
}
