package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"loki/pkg/app"
	"loki/pkg/e"
	"net/http"
)

func JWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		var (
			token string
			ecode = e.SUCCESS
		)
		if s, exist := c.GetQuery("token"); exist{
			token = s
		}else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = e.INVALID_PARAMS
		} else {
			_, err := app.ParseToken(token)
			if err !=nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					ecode = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
				log.Println("ParseToken failed!!", err)
			}
		}
		if ecode != e.SUCCESS{
			c.JSON(http.StatusOK, gin.H{
				"code": ecode,
				"msg": e.GetMsg(ecode),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
