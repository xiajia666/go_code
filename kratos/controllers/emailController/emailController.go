package emailController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"model/kratos/internal/pkg"
	"net/http"
	"net/smtp"
	"strings"
)

func EmailSend(ctx *gin.Context) {
	// 读取邮件服务配置文件
	config, _ := pkg.ReadYamlConfig("../../../", "config.yaml")
	emailSend, emailPassword := fmt.Sprintf("%v", config.Get("email.emailSend")), fmt.Sprintf("%v", config.Get("email.emailPassword"))

	name, emailAccept := ctx.Query("name"), ctx.Query("email")

	//设置发送方和接收方的邮箱
	e := email.NewEmail()
	e.From = fmt.Sprintf("dj <%s>", emailSend)
	e.To = []string{emailAccept}

	//设置主题
	e.Subject = "浙江工业大学C210实验室邀请函"

	//设置文件发送的内容
	html, _ := ioutil.ReadFile("../../html/send-Invitation.html")
	html_final := strings.ReplaceAll(string(html), "xx", name)
	e.Text = []byte(html_final)

	//设置附件
	e.AttachFile("../../images/person/Go语言中文文档.pdf")

	//设置服务器相关的配置
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", emailSend, emailPassword, "smtp.163.com"))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "message": "发送成功"})
}
