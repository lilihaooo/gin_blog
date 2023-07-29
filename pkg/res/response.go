package res

import (
	"blog_gin/global"
	"blog_gin/pkg/constant/res_const"
	"github.com/gin-gonic/gin"

	"net/http"
)

type response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Fail(c *gin.Context, code int, msg string) {
	message := (*global.ErrMap)[code]
	if msg != "" {
		message = msg
	}
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  message,
	})
}

func Ok(c *gin.Context) {
	code := res_const.SUCCESS
	msg := (*global.ErrMap)[code]
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
	})
}

func OkWithData(c *gin.Context, data interface{}) {
	code := res_const.SUCCESS
	msg := (*global.ErrMap)[code]
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
