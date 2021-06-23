package impl

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"sync"
	"time"
)

type UserAuthServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserAuthServiceImpl) GetLoginCode(username string) error {
	//TODO 检验账号是否合法
	//生成六位随机验证码发送
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	logging.Logger.Debug(code)
	//将验证码通过邮件发送给用户
	mail := email.NewEmail()
	//设置发送方的邮箱
	mail.From = common.Sender
	// 设置接收方的邮箱
	mail.To = []string{username}
	//设置主题
	mail.Subject = common.EmailSubject
	//设置文件发送的内容
	mail.Text = []byte("您的验证码为 " + code + " 有效期15分钟，请不要告诉他人哦！")
	//设置服务器相关的配置
	err := mail.Send(common.EmailServerAddr, smtp.PlainAuth("", common.Sender, common.EmailAuthorizationCode, common.EmailHost))
	if err != nil {
		logging.Logger.Error(err)
		return err
	}
	// 将验证码存入redis，设置过期时间为15分钟
	common.RedisUtil.SetEx(common.CodeKey+username, code, common.CodeExpireTime, time.Second)
	return nil
}
