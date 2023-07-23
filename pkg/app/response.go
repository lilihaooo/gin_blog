package app

import (
	"blog_gin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, msg string, data interface{}) {
	message := (*global.ErrMap)[code]
	if msg != "" {
		message = msg
	}
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
}
