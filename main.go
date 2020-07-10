package main

import (
	"github.com/gin-gonic/gin"
	"log"
	. "loki/routers"
	v1 "loki/routers/v1"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	server := InitWsServer()
	go server.Serve()
	defer server.Close()

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.TimeNow())
	apiv1.Use(Cors())
	{
		// ping
		apiv1.GET("/ping", v1.Ping)
		apiv1.GET("/ws", v1.Ws)
		apiv1.GET("/socket.io/*any", gin.WrapH(server))
		apiv1.POST("/socket.io/*any", gin.WrapH(server))
	}
	//r.Use(Cors())
	r.GET("/ws", v1.Ws)
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	//r.StaticFS("/public", http.Dir("/Users/xuchu/xcgo/loki/asset"))
	err := r.Run(":10900")
	if err != nil {
		log.Println(err)
	}
}
