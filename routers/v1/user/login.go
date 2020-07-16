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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	code := e.SUCCESS
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	loginPassword := user.Password
	//判断提交的用户名是否存在
	err = global.DBEngine.
		Where("username = ?", user.Username).
		First(&user).
		Error
	if err != nil {
		log.Println(err)
		code := e.LOGIN_FAILED
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
	if passwordIsOk {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.MsgFlags[code],
		})
		return
	}
	code = e.LOGIN_FAILED
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
	})
	return
}
