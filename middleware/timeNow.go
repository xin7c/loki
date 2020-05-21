package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func TimeNow() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeNowStr := time.Now().Unix()
		c.Set("timeNowStr", timeNowStr)
		c.Next()
	}
}
