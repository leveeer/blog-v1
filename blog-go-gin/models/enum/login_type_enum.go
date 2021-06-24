package enum

type LoginTypeEnum int

const (
	EMAIL LoginTypeEnum = iota
	QQ
	WEIBO
)

var (
	LoginTypeDesc = map[LoginTypeEnum]string{
		EMAIL: "邮箱登录",
		QQ:    "QQ登录",
		WEIBO: "微博登录",
	}

	LoginType = map[LoginTypeEnum]int8{
		EMAIL: 0,
		QQ:    1,
		WEIBO: 2,
	}
)

func (t LoginTypeEnum) GetLoginTypeDesc() string {
	return LoginTypeDesc[t]
}

func (t LoginTypeEnum) GetLoginType() int8 {
	return LoginType[t]
}
