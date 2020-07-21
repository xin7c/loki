package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/app"
	"loki/pkg/e"
	"loki/pkg/util"
	"net/http"
)

func Login(c *gin.Context) {
	code := e.SUCCESS
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	loginPassword := user.Password
	log.Println(user)
	//判断提交的用户名是否存在
	err = user.CheckUserExist(global.DBEngine)
	if err != nil {
		log.Printf("用户名不存在: %s", err)
		code := e.USER_NOT_FOUND
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.MsgFlags[code],
		})
		return
	}
	log.Println(user.Password)
	// 密码验证
	passwordIsOk := util.ValidatePassword(user.Password, loginPassword)
	log.Println("用户输入密码验证", passwordIsOk)
	// 验证密码通过
	if passwordIsOk {
		token, err := app.GenerateToken(user.Username, user.Password)
		if err != nil {
			code = e.ERROR
			log.Printf("passwordIsOk - 生成token err: %s", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":     code,
			"msg":      e.MsgFlags[code],
			"username": user.Username,
			"token":    token,
		})
		return
	}
	// 判断密码不匹配
	code = e.LOGIN_FAILED
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
	})
	return
}
