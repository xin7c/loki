package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/e"
	"net/http"
)

func Ping(c *gin.Context) {
	timeNowStr, _ := c.Get("timeNowStr")
	// model.Auth的其他字段需要编写gorm回调函数进行处理
	auth := model.Auth{Username: "777", Password: "777mima"}
	err := auth.Create(global.DBEngine)
	code := e.ERROR
	if err != nil {
		log.Println(err)
	}
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message":  "pong",
		"time": timeNowStr,
	})
	return
}
