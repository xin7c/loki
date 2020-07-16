package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/e"
	"net/http"
)

type ExistingUser struct {
	Username string `json:"username"`
	UserType string `json:"user_type"`
}

type ExistingUsers []ExistingUser

func GetUsers(c *gin.Context) {
	code := e.ERROR
	var users model.Users
	if err := global.DBEngine.Find(&users).Error; err != nil {
		log.Println(err)
	}
	eus := ExistingUsers{}
	for _, v := range users {
		eu := ExistingUser{
			Username: v.Username,
			UserType: v.Usertype.String,
		}
		eus = append(eus, eu)
	}
	log.Println(eus)
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  eus,
	})
	return
}
