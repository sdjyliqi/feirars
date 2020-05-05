package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/router"
)

func main() {
	flag.Parse()
	glog.Flush()
	r := gin.Default()
	//r.Use(middleware.RequestAddIPLoc())
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(":8899")

}
