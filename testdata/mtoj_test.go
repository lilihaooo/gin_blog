package testdata

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	FAIL_ADD_DATA  = 800

	//用户相关
	ERROR_EXIST_USER     = 10001
	ERROR_NOT_EXIST_USER = 10002
	ERROR_PASS_USER      = 10003
	ERROR_CAPTCHA_USER   = 10004
	FAIL_LOGOUT_USER     = 10005

	//token相关
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_CHECK_FAIL          = 20005

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

func TestMtoj(t *testing.T) {
	var MsgFlags = map[int]string{
		SUCCESS:        "ok",
		ERROR:          "fail",
		INVALID_PARAMS: "请求参数错误",
		FAIL_ADD_DATA:  "操作数据失败",

		//用户相关
		ERROR_EXIST_USER:     "已存在该用户名称",
		ERROR_NOT_EXIST_USER: "该用户不存在",
		ERROR_PASS_USER:      "用户密码错误",
		ERROR_CAPTCHA_USER:   "验证码错误",
		FAIL_LOGOUT_USER:     "退出失败",
		//token相关
		ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
		ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
		ERROR_AUTH_TOKEN:               "Token生成失败",
		ERROR_AUTH:                     "Token错误",
		ERROR_AUTH_CHECK_FAIL:          "无权限，请联系管理员",

		ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
		ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
		ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	}

	jsonData, err := json.Marshal(MsgFlags)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	// 将字节切片转换为字符串并输出
	fmt.Println(string(jsonData))
}
