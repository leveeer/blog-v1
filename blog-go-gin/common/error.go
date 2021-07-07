package common

import (
	"errors"
)

type ErrorCode uint

// user group
const (
	AccountIsBanned ErrorCode = iota + 10001
	UserNotFound
	CheckPasswordFail
	AdminPrivilegeNeeded
	GetRoomIdFail
	GetUserMatchGroupFail
	AlreadyInTeamFail
	NoGameServerFail
	ServerException
)

// request group
const (
	InvalidRequestParams ErrorCode = iota + 20001
	TokenCreateFail
	InvalidUserNameLength
	InvalidUserEmail
	CreateUserFail
	VerfiyUserFail
	AccountOrPhoneNumberDup
	SendEmailFail
	EmailIntervalTooShort
	VerifyCodeExpired
	CreatePayeeFail
	NotifyBillFail
	NotifyPayFail
	BanlancePayFail
	UnfinishedOrder
	InvalidAmount
	GetAssetRecordFail
	GetUserInfoFail
	GetUserPayOrderFail
	InvalidPhoneNum
	SMSIntervalTooShort
	SMSDayLimitExceed
	SMSSendFail
	ResetPassWordFail
	InvalidBindPhone
	InvalidPhone
	RecordNotFound
	GetOrderLogFail
	OrderNoNull
	NotEnoughMoney
	ModifyUserAssetFail
	GetMoneyChangeFail
	ModifyUserPayeeInfoFail
	ContactAdmin
	GetUserPayeeInfoFail
	UnKnowHolder
	UnKnowGameServer
	CreateOrderFail
	UpdateCoinFail
	GetUserBalanceFail
	UserNotEnoughMoney
	SendGiftFail
	RefreshFriendFail
	SearchRelationFail
	HasFollowUser
	InvalidGift
	GetArticlesFail
	GetBlogHomeInfoFail
	ApiCallTimeout
	GetArticleByIdFail
	GetArticleArchivesFail
	GetCategoriesFail
	GetTagsFail
	GetArticleByCategoryIDFail
	GetAboutFail
	GetMessagesFail
	AddMessageFail
	GetFriendLinksFail
	GetCommentsFail
	GetLoginCodeFail
	RegisterFail
	LoginFail
	AddCommentFail
	AddReplyFail
	GetRepliesFail
	GetUserMenusFail
)

// unknown group
const (
	UnknownErr ErrorCode = iota + 90001
	NoneErr
)

var Error = map[ErrorCode]error{
	// user group
	AccountIsBanned:       errors.New("账号已被封禁"),
	InvalidUserNameLength: errors.New("非法邮箱"),
	InvalidUserEmail:      errors.New("非法邮箱格式"),
	UserNotFound:          errors.New("用户不存在"),
	AdminPrivilegeNeeded:  errors.New("权限不足"),
	ContactAdmin:          errors.New("系统错误，请联系客服"),
	GetRoomIdFail:         errors.New("获取房间失败或该房间不是公开房间"),
	GetUserMatchGroupFail: errors.New("获取用户匹配倾向失败"),
	// unknown group
	UnknownErr:                 errors.New("未知错误"),
	InvalidRequestParams:       errors.New("参数非法"),
	CreateUserFail:             errors.New("创建用户失败"),
	VerfiyUserFail:             errors.New("验证用户失败"),
	AccountOrPhoneNumberDup:    errors.New("账号已注册或者手机号已被绑定"),
	SendEmailFail:              errors.New("发送邮件验证码失败"),
	EmailIntervalTooShort:      errors.New("请求验证码过于频繁"),
	VerifyCodeExpired:          errors.New("验证码已过期"),
	CreatePayeeFail:            errors.New("创建收款信息失败"),
	NotifyBillFail:             errors.New("代收回调失败"),
	NotifyPayFail:              errors.New("代付回调失败"),
	BanlancePayFail:            errors.New("系统发生错误，请联系管理员"),
	UnfinishedOrder:            errors.New("每次只能进行一个订单，请到订单列表完成或取消当前订单。"),
	InvalidAmount:              errors.New("非法金额"),
	GetAssetRecordFail:         errors.New("获取资金记录失败"),
	GetUserInfoFail:            errors.New("获取用户信息失败"),
	GetUserPayOrderFail:        errors.New("获取用户账单失败"),
	InvalidPhoneNum:            errors.New("不是有效的手机号"),
	SMSIntervalTooShort:        errors.New("请求验证码过于频繁"),
	SMSDayLimitExceed:          errors.New("验证码请求已达到每日限制"),
	SMSSendFail:                errors.New("发送验证码失败"),
	ResetPassWordFail:          errors.New("重置密码失败"),
	InvalidBindPhone:           errors.New("该手机已经绑定"),
	InvalidPhone:               errors.New("手机号不属于该用户"),
	CheckPasswordFail:          errors.New("密码错误,如为旧版邮箱账号请重新设置密码"),
	RecordNotFound:             errors.New("未找到订单"),
	GetOrderLogFail:            errors.New("获取订单日志失败"),
	OrderNoNull:                errors.New("订单号为空"),
	NotEnoughMoney:             errors.New("余额不足"),
	ModifyUserAssetFail:        errors.New("修改用户余额失败"),
	GetMoneyChangeFail:         errors.New("获取流水记录失败"),
	ModifyUserPayeeInfoFail:    errors.New("修改用户默认联系人失败"),
	GetUserPayeeInfoFail:       errors.New("获取用户支付信息失败"),
	TokenCreateFail:            errors.New("登录令牌创建错误"),
	AlreadyInTeamFail:          errors.New("你在队伍中"),
	NoGameServerFail:           errors.New("服务器维护中"),
	UnKnowHolder:               errors.New("获取服务器失败"),
	UnKnowGameServer:           errors.New("未知游戏服务器"),
	ServerException:            errors.New("系统异常，请稍后重试"),
	CreateOrderFail:            errors.New("生成订单失败"),
	UpdateCoinFail:             errors.New("扣除金币失败"),
	GetUserBalanceFail:         errors.New("获取用户余额失败"),
	UserNotEnoughMoney:         errors.New("用户剩余金币不足"),
	SendGiftFail:               errors.New("赠送礼物失败"),
	RefreshFriendFail:          errors.New("添加好友成功,但刷新房间好友列表失败"),
	SearchRelationFail:         errors.New("查找好友关系失败"),
	HasFollowUser:              errors.New("已经关注该用户"),
	InvalidGift:                errors.New("未知礼物类型"),
	GetArticlesFail:            errors.New("获取文章列表失败"),
	GetBlogHomeInfoFail:        errors.New("获取博客首页信息失败"),
	ApiCallTimeout:             errors.New("接口调用超时"),
	GetArticleByIdFail:         errors.New("获取文章失败"),
	GetArticleArchivesFail:     errors.New("获取归档失败"),
	GetCategoriesFail:          errors.New("获取分类失败"),
	GetTagsFail:                errors.New("获取标签失败"),
	GetArticleByCategoryIDFail: errors.New("获取分类文章失败"),
	GetAboutFail:               errors.New("获取关于失败"),
	GetMessagesFail:            errors.New("获取留言失败"),
	AddMessageFail:             errors.New("留言失败"),
	GetFriendLinksFail:         errors.New("获取友链失败"),
	GetCommentsFail:            errors.New("获取评论失败"),
	GetLoginCodeFail:           errors.New("获取验证码失败"),
	RegisterFail:               errors.New("注册失败"),
	LoginFail:                  errors.New("用户名或密码错误"),
	AddCommentFail:             errors.New("评论失败"),
	AddReplyFail:               errors.New("回复失败"),
	GetRepliesFail:             errors.New("获取回复列表失败"),
	GetUserMenusFail:           errors.New("获取用户菜单失败"),
}

func GetMsg(code ErrorCode) string {
	msg, ok := Error[code]
	if ok {
		return msg.Error()
	}

	return Error[UnknownErr].Error()
}
