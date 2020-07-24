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

type ModifyUser struct {
	model.User
	NewPassword string `json:"new_password"`
}

func Modify(c *gin.Context) {
	code := e.SUCCESS
	var modifyUser ModifyUser
	err := c.ShouldBindJSON(&modifyUser)
	if err != nil {
		code = e.INVALID_PARAMS
	}
	loginPassword := modifyUser.Password
	log.Println(modifyUser)
	//判断提交的用户名是否存在
	err = modifyUser.CheckUserExist(global.DBEngine)
	if err != nil {
		log.Printf("用户名不存在: %s", err)
		code := e.USER_NOT_FOUND
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": e.MsgFlags[code],
		})
		return
	}
	log.Println(modifyUser.Password, modifyUser.NewPassword)
	// 密码验证
	passwordIsOk := util.ValidatePassword(modifyUser.Password, loginPassword)
	log.Println("用户输入密码验证", passwordIsOk)
	// 验证密码通过
	if passwordIsOk {
		// 用户密码加密
		encodePWD, encodePasswordError := util.EncodePassword(modifyUser.NewPassword)
		if encodePasswordError != nil {
			code = e.ERROR
			log.Printf("/modify %s", encodePasswordError)
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.MsgFlags[code],
			})
			return
		}
		log.Printf("新密码加密结果：%s", encodePWD)
		err = modifyUser.Modify(global.DBEngine, encodePWD)
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": e.MsgFlags[code],
			"data":    "修改密码成功",
		})
		return
	}
	// 判断密码不匹配
	code = e.MODIFY_VERIFY_PASSWORD_FAILED
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.MsgFlags[code],
	})
	return

}
