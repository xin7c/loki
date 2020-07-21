package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//log.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,"+
			"X-CSRF-Token, Authorization,"+
			"Token, X-Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("content-type", "application/json;charset=UTF-8")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func InitRouter(mode string) *gin.Engine {
	r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	gin.SetMode(mode)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	//apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.TimeNow())
	//apiv1.Use(Cors())
	//{
	//	// ping
	//	apiv1.GET("/ping", v1.Ping)
	//	apiv1.GET("/ws", v1.Ws)
	//}
	//r.Use(Cors())
	r.GET("/ws/*any", gin.WrapH(server))
	r.StaticFS("/p", http.Dir("../asset"))
	return r
}
