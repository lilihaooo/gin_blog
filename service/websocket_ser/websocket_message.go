package websocket_ser

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/pkg/res"
	"blog_gin/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

var ws = make(map[uint]*websocket.Conn)

func (WebsocketService) WebsocketChat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.Fail(c, 32323, err.Error())
		global.Logrus.Error("Upgrade Websocket服务失败:", err)
		return
	}
	defer conn.Close()
	// 获取用户Id
	user, _ := c.Get("user")
	payload := user.(*jwts.Payload)
	// 将连接加入到map中
	ws[payload.UserID] = conn
	for {
		ms := message{}
		err := conn.ReadJSON(&ms)
		if err != nil {
			global.Logrus.Error("读取消息失败:", err)
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
			// 获取map中的值, 返回两次参数 值, 是否存在(bool)
			if conn, exists := ws[re.UserID]; exists {
				err = conn.WriteMessage(websocket.TextMessage, byteSlice) // 文本类型
				if err != nil {
					global.Logrus.Error("Write消息失败:", err)
					break
				}
			}
		}

	}
}
