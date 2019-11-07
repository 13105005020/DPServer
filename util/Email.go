package util

import (
	"net/smtp"
	"strings"
)

//发送邮件
//title 邮件标题
//body 邮件内容
//to 发送目标
func SendEmail(title, body string, to []string) {
	// 邮箱发送地址
	UserEmail := "1148031762@qq.com"
	MailSmtpPort := ":587"
	// Smtp密码
	MailPassword := "xztouumrrzvhiccc"
	// Smtp服务器
	MailSmtpHost := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, MailPassword, MailSmtpHost)
	nickname := "DP助手"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + UserEmail + ">\r\nSubject: " + title + "\r\n" + contentType + "\r\n\r\n" + body)
	for _, v := range SliceSlice(to, 1) {
		smtp.SendMail(MailSmtpHost+MailSmtpPort, auth, UserEmail, v, msg)
	}
}
