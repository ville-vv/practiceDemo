package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

//发送邮件的逻辑函数
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	// 邮箱账号
	user := "xxxx@163.com"
	//注意，此处为授权码、不是密码
	password := "123456"
	//smtp地址及端口
	host := "smtp.163.com:25"
	//接收者，内容可重复，邮箱之间用；隔开
	to := "xxxx@163.com"
	//邮件主题
	subject := "测试通过golang发送邮件"
	//邮件内容
	text := "你好！"
	body := `
    <html>
    <body>
    <h3>
    "测试通过golang发送邮件"` + text + `
    </h3>
    </body>
    </html>
    `
	fmt.Println("send email")
	//执行逻辑函数
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!")
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!")
	}

}