package middlewares

import (
	"time"

	"github.com/baaj2109/shorturl/common"
	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := common.GetUUID()
		startTime := time.Now()
		common.SugarLogger.Info("request start: %s %s", id, startTime)
		c.Next()
		now := time.Now().Format("2006-01-02 15:04:05.000")
		duration := int(time.Now().Sub(startTime) / 1e6) //单位毫秒
		request := c.Request.RequestURI
		host := c.Request.Host
		clientIp := c.ClientIP()
		code := c.Writer.Status()
		ua := c.Request.UserAgent()
		common.SugarLogger.Info("request_end: %s %s %d %s %s %s %d %s", id, now, duration, request, host, clientIp, code, ua)

	}
}
