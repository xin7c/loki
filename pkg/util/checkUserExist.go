package util

import (
	"log"
	"loki/global"
	"loki/internal/model"
)

func CheckUserExist(username string) bool {
	var user model.User
	notFound := global.DBEngine.Where("username = ?", username).First(&user).RecordNotFound()
	if notFound == true {
		log.Println("不存在")
		return false
	}
	return true
}
