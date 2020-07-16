package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncodePassword(pwd string) (string, error) {
	//加密处理
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	encodePWD := string(hash)
	log.Printf("落库密码: %s", encodePWD)
	return encodePWD, nil
}
