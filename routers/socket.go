package routers

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/url"
	"time"
)

type Msg struct {
	UserId    string   `json:"userId"`
	Text      string   `json:"text"`
	State     string   `json:"state"`
	Namespace string   `json:"namespace"`
	Rooms     []string `json:"rooms"`
}

func InitWsServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		m, _ := url.ParseQuery(s.URL().RawQuery)
		s.SetContext(m["name"][0])
		s.Join("main")
		//log.Println(s.URL().RawQuery)
		//log.Println(m["name"])
		//log.Println(m["id"])
		log.Printf("连接事件: %s", s.ID())
		roomlist := server.Rooms("/")
		roomlen := server.RoomLen("/", s.ID())
		log.Printf("连接事件: 用户加入的房间号[%+#v]", s.Rooms())
		log.Printf("连接事件: 当前系统所有房间[%+#v]", roomlist)
		connMsg := map[string]interface{}{
			"msg":     fmt.Sprintf("conn:sid[%s]", s.ID()),
			"connid":  s.ID(),
			"roomlen": roomlen,
			"rooms":   s.Rooms(),
		}
		s.Emit("connected", connMsg)
		//s.Emit("conn", m)
		return nil
	})
	server.OnEvent("/", "join", func(s socketio.Conn, room string, username string) {
		// s.Rooms() 返回用户加入的房间号
		// server.Rooms() 则记录了当前系统所有的房间号
		s.Join(room)
		s.SetContext(username)
		log.Println(username + "加入房间@ " + room)
		roomlen := server.RoomLen("/", room)
		roomlist := server.Rooms("/")
		server.BroadcastToRoom("/", room, "join",
			map[string]interface{}{
				"msg":      fmt.Sprintf("[%s]加入房间%s[@%d]", username, room, time.Now().Unix()),
				"roomid":   room,
				"roomlen":  roomlen,
				"roomlist": roomlist, //s.Rooms(),
			})
		//msg := Msg{s.ID(), "<= " + s.ID() + " join " + room, "state", s.Namespace(), s.Rooms()}
		//for k,v := range server.Rooms("/"){
		//	log.Println(k, v)
		//}
		log.Printf("房间名[%s]-namespace[%s]-username[%s]-房间人数[%d]",
			room, s.Namespace(), username, roomlen)
		log.Printf("用户加入的房间号[%+#v]", s.Rooms())
		log.Printf("当前系统所有房间[%+#v]", roomlist)
	})

	server.OnEvent("/", "leave", func(s socketio.Conn, room string) {
		//roomlist := server.Rooms("/")
		log.Println(s.Context().(string) + "离开房间@ " + room)
		s.Leave(room)
		reGetRoomlist := server.Rooms("/")
		s.Emit("leave", fmt.Sprintf("你已从[%s]房间离开,系统存活房间数:[%+#v]", room, reGetRoomlist))
		log.Printf("用户加入的房间号[%+#v]", s.Rooms())
		log.Printf("当前系统所有房间[%+#v]", reGetRoomlist)

	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg, room, username string) string {
		//s.SetContext(msg)
		//log.Println("/ notice => ", s.URL().RawQuery)
		log.Printf("[聊天信息]%s说: %s @ %s", username, msg, room)
		//log.Println("notice s.Context().(string)", s.Context().(string))
		//s.Emit("chat", "chat: " + msg + " SID:/ => chat: " + s.ID())
		server.BroadcastToRoom("/", room, "chat",
			map[string]interface{}{
				"msg": fmt.Sprintf("[%s]在房间%s[@%d]说: %s", username, room, time.Now().Unix(), msg),
			})
		return "recv " + msg
	})
	server.OnEvent("/", "joinRoom", func(s socketio.Conn, msg string) string {
		//s.SetContext(msg)
		//log.Println("/ notice => ", s.URL().RawQuery)
		log.Println("/ joinRoom =>", msg, s.ID())
		//log.Println("notice s.Context().(string)", s.Context().(string))
		s.Emit("joinRoom", "/ => joinRoom "+" SID:/ => notice: "+s.ID())
		return "recv "
	})
	//server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	//	log.Println("notice:", msg)
	//	s.Emit("reply", "have " + msg)
	//})
	//server.OnEvent("/", "bye", func(s socketio.Conn) string {
	//	last := s.Context().(string)
	//	s.Emit("bye", last)
	//	s.Close()
	//	return last
	//})
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	//
	//server.OnConnect("/chat", func(s socketio.Conn) error {
	//	s.SetContext("")
	//	log.Println(s.URL().RawQuery)
	//	log.Println("chat => connected:", s.ID())
	//	s.Emit("chat => conn", "conn : sid " + s.ID())
	//	return nil
	//})
	server.OnEvent("/chat", "chat", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		log.Println("chat => ", s.URL().RawQuery)
		log.Println("chat => notice:", msg)
		s.Emit("chat", "chat recv "+msg)
		return "recv " + msg
	})
	//server.OnError("/chat", func(s socketio.Conn, e error) {
	//	log.Println("chat => meet error:", e)
	//})
	//server.OnDisconnect("/chat", func(s socketio.Conn, reason string) {
	//	log.Println("chat => closed", reason)
	//})
	return server
}
