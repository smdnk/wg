package demo

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 将http协议升级为 websocket协议
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		// 读取消息
		messageType, msgText, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(messageType, string(msgText))
		// 发送消息
		_ = conn.WriteMessage(websocket.TextMessage, msgText)
	}

}
func StartSocket() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	http.HandleFunc("/ws", handleWebSocket)
	_ = http.ListenAndServe(":8080", nil)

	wg.Wait()
}
