package res

import (
	"blog_gin/global"
	"github.com/gin-gonic/gin"

	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type listResponse[T any] struct {
	List  T     `json:"list"`
	Count int64 `json:"count"`
}

// ValidateResponse 验证失败响应
type ValidateResponse struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func Fail(c *gin.Context, code int, msg string) {
	message := (*global.ResMap)[code]
	if msg != "" {
		message = msg
	}
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  message,
	})
}

func FailValidate(c *gin.Context, data []ValidateResponse) {
	message := (*global.ResMap)[INVALID_PARAMS]
	c.JSON(http.StatusOK, response{
		Code: INVALID_PARAMS,
		Msg:  message,
		Data: data,
	})
}

func Ok(c *gin.Context) {
	code := SUCCESS
	msg := (*global.ResMap)[code]
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
	})
}

func OkWithMsg(c *gin.Context, msg string) {
	code := SUCCESS
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
	})
}
func OkWithData(c *gin.Context, data interface{}) {
	code := SUCCESS
	msg := (*global.ResMap)[code]
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OkWithList(c *gin.Context, list any, count int64) {
	OkWithData(c, listResponse[any]{
		List:  list,
		Count: count,
	})
}
