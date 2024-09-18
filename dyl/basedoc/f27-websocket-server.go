// websocket是socket连接和http协议的结合体，可以实现网页和服务端的长连接
// go get github.com/gorilla/websocket  安装  Go 环境中用于获取和安装 gorilla/websocket 包，它是一个用来处理 WebSocket 协议的库。
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connLis []*websocket.Conn

func handler(res http.ResponseWriter, req *http.Request) {
	// 服务升级
	conn, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	connLis = append(connLis, conn)
	for {
		// 消息类型，消息，错误
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		//从客户都拿到p的值。
		//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))

		for index := range connLis {
			//服务端给多个客户端转发。
			//这里有点类似，php的workman.基于这个，其实可以做一个简易的聊天室。
			connLis[index].WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
		}

		//打印输出p
		fmt.Println(t, string(p))
	}
	defer conn.Close()
	fmt.Println("服务关闭")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
