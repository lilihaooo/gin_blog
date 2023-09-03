package websocket_ser

import "github.com/gorilla/websocket"

type WebsocketService struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	Data   string `json:"data"`
	RoomID uint   `json:"room_id"`
}
