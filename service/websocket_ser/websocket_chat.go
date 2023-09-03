package websocket_ser

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

type message struct {
	Data   string `json:"data"`
	RoomID uint   `json:"room_id"`
}

var wsChatMap = make(map[uint]*websocket.Conn)

func (WebsocketService) WebsocketChat(c *gin.Context) {
	// 将http升级为websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logrus.Error("Upgrade Websocket服务失败:", err)
		return
	}
	user := c.MustGet("user")
	payload := user.(*jwts.Payload)
	wsChatMap[payload.UserID] = conn
	defer func() {
		conn.Close()
		delete(wsChatMap, payload.UserID)
	}()
	// 启动心跳检测
	go handleHeartbeat(conn)

	for {
		ms := message{}
		err := conn.ReadJSON(&ms)
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				global.Logrus.Error("读取消息失败:", err)
			}
			break
		}
		// 获取消息之后用户是否属于这个房间
		var userRoomModel models.UserRoomModel
		in := userRoomModel.UserInRoom(payload.UserID, ms.RoomID)
		if !in {
			global.Logrus.Error("该用户不在此房间")
			break
		}
		// 获取该房间里的在线用户
		condition := "room_id = " + strconv.Itoa(int(ms.RoomID))
		records, err := userRoomModel.GetRecordsWithCondition(condition)
		if err != nil {
			global.Logrus.Error("查询数据库失败", err)
			break
		}
		// 将消息保存到数据库
		var messageModel models.MessageModel
		err = messageModel.AddOneMessage(&models.MessageModel{
			UserID: payload.UserID,
			RoomID: ms.RoomID,
			Data:   ms.Data,
		})
		if err != nil {
			global.Logrus.Error("数添加消息记录失败", err)
			break
		}
		str := ms.Data
		byteSlice := []byte(str)
		// 发消息给相应的连接
		for _, re := range records {
			// 获取map中的值, 返回两次参数值, 是否存在(bool)
			if conn, exists := wsChatMap[re.UserID]; exists {
				err = conn.WriteMessage(websocket.TextMessage, byteSlice) // 文本类型
				if err != nil {
					global.Logrus.Error("Write消息失败:", err)
					break
				}
			}
		}
	}
}

func handleHeartbeat(conn *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second) // 每10秒发送一次心跳消息
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
				return
			}
		}
	}
}
