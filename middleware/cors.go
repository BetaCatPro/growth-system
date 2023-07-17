package middleware

import (
	"github.com/gin-gonic/gin"

	"growth/comm"
	"growth/global"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// prometheus 指标
		comm.MetricAdd()

		// 支持跨域
		origin := c.GetHeader("Origin")
		if global.AllowOrigin[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTION")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		c.Next()
	}
}
