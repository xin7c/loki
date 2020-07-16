package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password, loginPassword string) bool {
	// 密码验证
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPassword)) //验证（对比）
	if err != nil {
		fmt.Println("pwd wrong")
		return false
	} else {
		fmt.Println("pwd ok")
		return true
	}
}
