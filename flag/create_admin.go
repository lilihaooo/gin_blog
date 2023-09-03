package flag

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/utils"
	"fmt"
)

type AdminCreateRequest struct {
	NickName   string ` validate:"required,max=36"`
	UserName   string `validate:"required,max=36"`
	Password   string `validate:"required,max=6"`
	RePassword string `validate:"required,max=6"`
	Email      string
}

func CreateAdmin() {
	var cr AdminCreateRequest
	fmt.Print("请输入用户名: ")
	fmt.Scan(&cr.UserName)
	fmt.Print("请输入昵称: ")
	fmt.Scan(&cr.NickName)
	fmt.Print("请输入密码: ")
	fmt.Scan(&cr.Password)
	fmt.Print("请再次输入密码: ")
	fmt.Scan(&cr.RePassword)
	if cr.Password != cr.RePassword {
		fmt.Println("两次输入的密码不一致!!!")
		return
	}
	fmt.Print("请输入用邮箱: ")
	fmt.Scanln(&cr.Email) // Scanln 可以不输入

	// 验证参数
	vRes := utils.ZhValidate(&cr)
	if vRes != nil {
		fmt.Printf("err: %s!!! \n", vRes[0].Msg)
		return
	}
	// 判断用户是否存在
	var userModel models.UserModel
	if row := global.DB.Take(userModel, "user_name = ?", cr.UserName).RowsAffected; row > 0 {
		fmt.Println("用户已存在!!!")
		return
	}

	// 对密码进行加密处理
	hashPassword, err := utils.HashPassword(cr.Password)
	if err != nil {
		fmt.Println("密码加密失败!!!")
		return
	}
	userModel.UserName = cr.UserName
	userModel.NickName = cr.NickName
	userModel.Password = hashPassword
	userModel.Avatar = "/uploads/avatar/default.jpg"
	userModel.Addr = "内网地址"
	userModel.IP = "127.0.0.1"
	userModel.Email = cr.Email
	userModel.Role = ctype.PermissionAdmin
	userModel.SignStatus = ctype.SignEmail // todo ?
	if err = global.DB.Create(&userModel).Error; err != nil {
		fmt.Println("添加失败!!!")
		return
	}
	fmt.Println("添加成功!!!")
}
