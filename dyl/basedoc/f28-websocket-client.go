package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go send(conn)
	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(t, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
