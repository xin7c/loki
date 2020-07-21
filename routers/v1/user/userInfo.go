package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/app"
	"loki/pkg/e"
	"loki/pkg/util"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var err error
	code := e.SUCCESS
	token := c.GetHeader("token")
	// 因为jwt中间件处理过了所以放心食用
	m, _ := app.ParseToken(token)
	log.Printf("解析token是否正确: %s - %s", m.AppKey, m.AppSecret)

	// 获取username并查询权限信息
	var user model.User
	err = c.ShouldBindJSON(&user)
	// log.Println("获取username并查询权限信息: ",user)
	if err != nil {
		log.Printf("获取username并查询权限信息 绑定结构体user: %s", err)
	}
	// 校验username与jwt中的app_key是否一致
	if !(util.EncodeMD5(user.Username) == m.AppKey) {
		code = e.USERINFO_CHECK_JWT_FAILED
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	err = user.GetUserInfo(global.DBEngine)
	if err != nil {
		log.Printf("userInfo.GetUserInfo(global.DBEngine): %s", err)
		code = e.USER_NOT_FOUND
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	// log.Println("user.GetUserInfo(global.DBEngine): ",user)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": map[string]interface{}{
			"roles":        user.Usertype.String,
			"introduction": fmt.Sprintf("I am %s", user.Username),
			"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"name":         user.Username,
		},
	})
	return
}
