package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/model"
	"loki/pkg/e"
	"net/http"
)

func Auths(c *gin.Context) {
	code := e.ERROR
	var as model.Auths
	if err := global.DBEngine.Find(&as).Error; err != nil {
		log.Println(err)
	}
	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": as,
	})
	return
}
