package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"net/http"
)

func Ping(c *gin.Context) {
	timeNowStr, _ := c.Get("timeNowStr")

	auth := model.Auth{Username: "777", Password: "777mima"}
	err := auth.Create(global.DBEngine)
	if err != nil{
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "pong",
		"time": timeNowStr,
	})
	return
}
