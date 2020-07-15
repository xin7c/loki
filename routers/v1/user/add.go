package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/e"
	"net/http"
)

func Add(c *gin.Context) {
	code := e.SUCCESS
	var userStruct model.User
	err := c.ShouldBindJSON(&userStruct)
	if err != nil {
		log.Println("c.BindJSON", err)
	}
	log.Printf("[%v,%v]", userStruct.Username, userStruct.Password)
	if userStruct.Username != "" && userStruct.Password != "" {
		user := model.User{Username: userStruct.Username, Password: userStruct.Password}
		err = user.Add(global.DBEngine)
		if err != nil {
			code = e.ERROR
			log.Println("user.Add", err)
		}
	} else {
		code = e.LOGIN_INVALID_PARAMS
		log.Println("参数错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
	})
	return
}
