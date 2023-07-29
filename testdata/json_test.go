package testdata

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {

	user := User{Password: "pwd", Email: "user@163.com"}
	//person := Person{23, "上海"}
	// 忽略掉 Password 字段
	data, _ := json.Marshal(struct {
		*User
		//Password string `json:"password,omitempty"`
	}{User: &user})
	fmt.Println("忽略字段: ", string(data)) // 打印: 忽略字段: {"u_name":"admin","email":"user@163.com"}

}

type User struct {
	Name     string `json:"u_name,omit(list)"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Person struct {
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}
