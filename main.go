package main

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/feirars/middleware"
	"github.com/sdjyliqi/feirars/router"
)

func main() {
	r := gin.Default()
	r.Use(middleware.RequestIDMiddleware())
	r.Use(ginprom.PromMiddleware(nil))
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(":8899")
}
