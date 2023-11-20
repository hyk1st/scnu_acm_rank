package middle

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"scnu_acm_rank/biz/config"
	"time"
)

type emailConfig struct {
	from     string
	subject  string
	password string
	host     string
}

var E emailConfig

func init() {
	E = emailConfig{}
	config.Add(&E)
}

func (e *emailConfig) Update() {
	e.host = config.Conf.EmailHost
	e.password = config.Conf.EmailPassword
	e.from = config.Conf.EmailFrom
	e.subject = config.Conf.EmailSubject
}

func SendEmail(to []string) error {
	e := email.NewEmail()
	e.From = E.from
	e.To = to
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%v", time.Now().UnixNano()%10)
	}
	e.Subject = E.subject
	e.Text = []byte("欢迎注册hyk online judge，您的验证码是： ")
	e.HTML = []byte("<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", E.from, E.password, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	AddCode(to[0], code)
	return nil
}
