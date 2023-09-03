package u_email

import (
	"blog_gin/global"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(subject, message, toEmail string) error {
	// 创建一个新的邮件对象
	e := email.NewEmail()
	// 设置发件人
	e.From = global.Config.Email.From
	// 设置收件人信息
	e.To = []string{toEmail}
	// 设置邮件主题和内容
	e.Subject = subject
	e.Text = []byte(message)
	// 使用 QQ 邮箱的 SMTP 服务器和端口
	host := global.Config.Email.Host
	port := global.Config.Email.Port
	password := global.Config.Email.Password
	return e.Send(fmt.Sprintf("%s:%d", host, port), smtp.PlainAuth("", e.From, password, host))
}
