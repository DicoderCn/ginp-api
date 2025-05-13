// Package emails
// @Author: zhangdi
// @File: email
// @Version: 1.0.0
// @Date: 2023/6/25 15:08
package email

import "gopkg.in/gomail.v2"

//const Email = "1778871523@qq.com"
//const Pwd = "rwehrxsnnvrneaae"

// EmailClient 定义邮件客户端结构体
type EmailClient struct {
	config EmailConfig
}

// SendEmail 发送邮件方法
func (ec *EmailClient) SendEmail(toEmail string, title string, content string) error {
	// 邮件配置
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", ec.config.Email)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", title)
	mailer.SetBody("text/html", content)

	// 发送邮件
	dialer := gomail.NewDialer(ec.config.Host, ec.config.Port, ec.config.Email, ec.config.Pwd)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	println("Email sent successfully!")

	return nil
}

// 删除原来的常量定义
// const Email = "ideatools@qq.com"
// const Pwd = "rtcxcomjggocbefb"
// const Port = 25
// const Host = "smtp.qq.com"

// 删除原来的 SendEmail1 函数
// func SendEmail1(toEmail string, title string, content string) error {
// 	// ...
// }
