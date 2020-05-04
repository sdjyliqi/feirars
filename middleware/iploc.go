package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/sdjyliqi/feirars/conf"
	"log"
	"net"
	"strings"
)

var Reader *geoip2.Reader

func init() {
	var err error
	Reader, err = geoip2.Open(conf.DefaultConfig.IPLOC)
	if err != nil {
		log.Fatalf("[init]Load GeoLite2-City.mmdb failed,err:%+v", err)
	}
}

//... RequestAddIPLoc  add request id into header
func RequestAddIPLoc() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		clientIP = "222.195.192.13"
		ip := net.ParseIP(clientIP)

		record, err := Reader.City(ip)
		if err == nil && record != nil && record.City.Names != nil && len(record.City.Names) > 0 {
			c.Request.Header.Add("IPLOC", strings.ToUpper(record.City.Names["en"]))
		}

		if err == nil && record != nil && record.Subdivisions != nil && len(record.Subdivisions) > 0 {
			province := record.Subdivisions[0].Names["en"]
			c.Request.Header.Add("IPPROVINCE", strings.ToUpper(province))
		}
		c.Next()
	}
}
