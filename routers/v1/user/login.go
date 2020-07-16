package user

import (
	"github.com/gin-gonic/gin"
	"loki/pkg/e"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	code := e.SUCCESS
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	//TODO 需要补充查库操作
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
	})
	return
}
