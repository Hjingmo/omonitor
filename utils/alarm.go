package utils

import (
	"fmt"
	"net/http"
	"net/smtp"
	"net/url"
	"time"

	"github.com/astaxie/beego"
)

func SendSms(mobiles []int64, msg string) {
	for _, mobile := range mobiles {
		if mobile < 10000000000 {
			break
		}
		go smsSend(mobile, msg)
	}
}

func smsSend(mobile int64, msg string) {
	smsOn, err := beego.AppConfig.Bool("sms")
	if err != nil {
		smsOn = false
	}
	if smsOn {
		Client := http.Client{Timeout: 2 * time.Second}
		smsUrl := fmt.Sprintf("http://sms.xxx.com:xxx%dxxx%sxxx", mobile, url.QueryEscape(msg))
		req, _ := http.NewRequest("GET", smsUrl, nil)
		resp, _ := Client.Do(req)
		defer resp.Body.Close()
	}
}

func SendEmail(toUsers []string, data string) {
	emailOn, err := beego.AppConfig.Bool("email")
	if err != nil {
		emailOn = false
	}
	if emailOn {
		emailUser := beego.AppConfig.String("email_user")
		emailPasswd := beego.AppConfig.String("email_passwd")
		emailHost := beego.AppConfig.String("email_host")
		auth := smtp.PlainAuth("", emailUser, emailPasswd, emailHost)
		msg := []byte("To: " + toUsers[0] + "\r\nFrom: +" + emailUser + ">\r\nSubject: 邮件报警\r\nContent-Type: text/html;charset=UTF-8\r\n\n " + data)
		err := smtp.SendMail(emailHost, auth, emailUser, toUsers, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
