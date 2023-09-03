package testdata

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
	"time"
)

// websocket的服务端, 客户端使用postman
func TestWebsocketServer(t *testing.T) {
	http.HandleFunc("/echo", echoHandler)
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



func echoHandler(w http.ResponseWriter, r *http.Request) {
	// 将http升级为websocket服务
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级websocket服务失败:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息失败:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)
		str := "今天天气正好"
		currentTime := time.Now()
		timeString := currentTime.Format("2006-01-02 15:04:05") // 使用自定义的时间格式模板

		byteSlice := []byte(str + timeString)

		err = conn.WriteMessage(messageType, byteSlice)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
