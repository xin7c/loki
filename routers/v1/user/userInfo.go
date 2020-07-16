package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"loki/pkg/app"
	"loki/pkg/e"
	"net/http"
)

func UserInfo(c *gin.Context) {
	code := e.SUCCESS
	token := c.GetHeader("token")
	if token == "" {
		code = e.INVALID_PARAMS
	} else {
		m, err := app.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
			log.Println("ParseToken failed!!", err)
		}
		log.Println(m.AppKey, m.AppSecret)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
	return
}
