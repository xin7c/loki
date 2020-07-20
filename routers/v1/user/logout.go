package user

import (
	"github.com/gin-gonic/gin"
	"loki/pkg/e"
	"net/http"
)

func Logout(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
		"data": "登出成功",
	})
	return
}
