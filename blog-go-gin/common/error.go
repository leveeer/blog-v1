package common

import (
	"errors"
)

type ErrorCode uint

// user group
const (
	AdminPrivilegeNeeded ErrorCode = iota + 10001
)

// request group
const (
	InvalidRequestParams ErrorCode = iota + 20001
	VerifyCodeExpired
	GetUserInfoFail
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
	GetHomeDataFail
	GetArticleOptionsFail
	UploadImageFail
	AddArticleFail
	TokenHasExpired
	UpdateArticleFail
	AddOrUpdateCategoryFail
	DeleteCategoryFail
	DeleteArticleFail
	AddOrUpdateTagFail
	DeleteTagFail
	LikeArticleFail
	LikeCommentFail
)

// unknown group
const (
	UnknownErr ErrorCode = iota + 90001
	NoneErr
)

var Error = map[ErrorCode]error{
	// user group
	AdminPrivilegeNeeded: errors.New("权限不足"),
	// unknown group
	UnknownErr:                 errors.New("未知错误"),
	InvalidRequestParams:       errors.New("参数非法"),
	VerifyCodeExpired:          errors.New("验证码已过期"),
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
	GetHomeDataFail:            errors.New("获取首页数据失败"),
	GetArticleOptionsFail:      errors.New("获取文章相关数据失败"),
	UploadImageFail:            errors.New("上传图片失败"),
	AddArticleFail:             errors.New("添加文章失败"),
	TokenHasExpired:            errors.New("token is "),
	UpdateArticleFail:          errors.New("更新文章失败"),
	AddOrUpdateCategoryFail:    errors.New("增加或更新分类失败"),
	AddOrUpdateTagFail:         errors.New("增加或更新标签失败"),
	DeleteCategoryFail:         errors.New("删除分类失败"),
	DeleteArticleFail:          errors.New("删除文章失败"),
	DeleteTagFail:              errors.New("删除标签失败"),
	LikeArticleFail:            errors.New("点赞文章失败"),
	LikeCommentFail:            errors.New("点赞评论失败"),
}

func GetMsg(code ErrorCode) string {
	msg, ok := Error[code]
	if ok {
		return msg.Error()
	}

	return Error[UnknownErr].Error()
}
