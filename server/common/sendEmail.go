package common

import (
	"2021/yunsongcailu/yunsong_server/config"
	"2021/yunsongcailu/yunsong_server/tools"
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmailCode(toEmail string) (theCode string,err error) {
	ec := config.Conf.Email
	fromEmail := ec.FromEmail
	fmt.Println(fromEmail)
	subject := ec.EmailSubject
	smtpAddr := ec.SmtpAddr
	smtpPort ,_ := strconv.Atoi(ec.SmtpPort)
	smtpPass := ec.SmtpPass
	// 生成验证码随机数
	code := tools.GetRandomCode()
	mailHeader := map[string][]string {
		"From":{fromEmail},
		"To":{toEmail},
		"Subject":{subject},
	}
	m:= gomail.NewMessage()
	m.SetHeaders(mailHeader)
	m.SetBody("text/html","尊敬的用户<br>您的注册验证码为:"+code+"<br>欢迎您!")
	d := gomail.NewDialer(smtpAddr,smtpPort,fromEmail,smtpPass)
	err = d.DialAndSend(m)
	return code,err
}