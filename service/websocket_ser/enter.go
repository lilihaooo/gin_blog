package websocket_ser

import (
	"github.com/gorilla/websocket"
)

type WebsocketService struct {
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
