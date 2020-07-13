package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/pkg/app"
	"loki/pkg/e"
	"net/http"
)

func GetAuth(c *gin.Context) {
	// 此处解析前端传递的参数，返回token
	code := e.ERROR
	t, err := app.GenerateToken("xuchu", "11")
	if err != nil{
		log.Println(err)
	}
	m, err := app.ParseToken(t)
	if err !=nil{
		log.Println("ParseToken failed!!", err)
	}
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"token":  t,
		"m": m,
	})
	return
}