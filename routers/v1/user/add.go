package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/e"
	"loki/pkg/util"
	"net/http"
)

func Add(c *gin.Context) {
	code := e.SUCCESS
	var userStruct model.User
	err := c.ShouldBindJSON(&userStruct)
	if err != nil {
		log.Printf("c.BindJSON解析失败 %s", err)
	}
	log.Printf("[%v,%v]", userStruct.Username, userStruct.Password)
	if userStruct.Username != "" && userStruct.Password != "" {
		// 判断用户名是否存在
		// 存在：true
		// 不存在：false
		userExist := util.CheckUserExist(userStruct.Username)
		if userExist {
			log.Printf("用户名%s已存在！", userStruct.Username)
			code = e.DUPLICATE_USERNAME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.MsgFlags[code],
			})
			return
		}
		// 用户密码加密
		encodePWD, encodePasswordError := util.EncodePassword(userStruct.Password)
		if encodePasswordError != nil {
			code = e.ERROR
			log.Printf("/add %s", encodePasswordError)
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.MsgFlags[code],
			})
			return
		}
		user := model.User{
			Username: userStruct.Username,
			Password: encodePWD,
		}
		err = user.Add(global.DBEngine)
		if err != nil {
			code = e.DUPLICATE_USERNAME
			log.Printf("/add %s", err)
		}
	} else {
		code = e.LOGIN_INVALID_PARAMS
		log.Println("参数错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.MsgFlags[code],
	})
	return
}
