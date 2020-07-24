package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"loki/global"
	"loki/internal/middleware"
	"loki/internal/model"
	"loki/pkg/setting"
	. "loki/routers"
	"loki/routers/v1"
	"loki/routers/v1/user"
	"time"
)

func setupSetting() error {
	lokiSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = lokiSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = lokiSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err = lokiSetting.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	//env := os.Getenv("ENV")
	//log.Printf("env: %s", env)
	err := setupSetting()
	if err != nil {
		log.Fatalf("setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("setupDBEngine err: %v", err)
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	gin.SetMode(global.ServerSetting.RunMode)
	server := InitWsServer()
	go server.Serve()
	defer server.Close()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.TimeNow())
	apiv1.Use(Cors(), middleware.JWT())
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.GET("/as", v1.Auths)
		apiv1.POST("/userinfo", user.GetUserInfo)
		apiv1.POST("/get_users", user.GetUsers)

	}
	r.Use(Cors())
	r.GET("/auth", v1.GetAuth)
	r.POST("/add", user.Add)
	r.POST("/login", user.Login)
	r.POST("/modify", user.Modify)
	r.GET("/logout", user.Logout)
	r.GET("/ws", v1.Ws)
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	//r.StaticFS("/public", http.Dir("/Users/xuchu/xcgo/loki/asset"))
	err := r.Run(fmt.Sprintf(":%s", global.ServerSetting.HttpPort))
	if err != nil {
		log.Println(err)
	}
}
