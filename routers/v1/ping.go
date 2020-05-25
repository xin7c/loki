package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	timeNowStr, _ := c.Get("timeNowStr")
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
		"time": timeNowStr,
	})
	return
}