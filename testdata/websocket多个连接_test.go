package testdata

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
	"time"
)

// websocket的服务端, 客户端使用postman
func TestWebsocketServer2(t *testing.T) {
	http.HandleFunc("/echo", echoHandler2)
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}

// 选择空的 struct{} 作为值类型是为了在映射中实现一种类似集合的结构，而不需要存储实际的数据。这样的结构通常被称为 "集合" 或 "集合映射"
var ws = make(map[*websocket.Conn]struct{})

func echoHandler2(w http.ResponseWriter, r *http.Request) {
	// 将http升级为websocket服务
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级websocket服务失败:", err)
		return
	}
	defer conn.Close()
	// 将连接加入到map中
	ws[conn] = struct{}{}

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

		// 将消息发到所有的连接中
		for conn := range ws {
			err = conn.WriteMessage(messageType, byteSlice)
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}

	}
}

// gin框架集成websocket
func TestGinWebsocketServer(t *testing.T) {
	r := gin.Default()
	r.GET("/echo2", Ws)
	r.GET("/hello", Ht)
	err := r.Run("127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
	}

}

func Ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级websocket服务失败:", err)
		return
	}
	defer conn.Close()
	// 将连接加入到map中
	ws[conn] = struct{}{}
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息失败:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)
		str := "gin今天天气正好"
		currentTime := time.Now()
		timeString := currentTime.Format("2006-01-02 15:04:05") // 使用自定义的时间格式模板
		byteSlice := []byte(str + timeString)

		// 将消息发到所有的连接中
		for conn := range ws {
			err = conn.WriteMessage(messageType, byteSlice)
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	}
}

func Ht(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mag": "test"})
}
