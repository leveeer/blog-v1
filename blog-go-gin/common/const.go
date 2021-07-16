package common

const (
	BloggerId              = 1 //博主id
	False                  = 0 //否
	True                   = 1 //是
	AppSecretKey           = "0qPhpSBdQXfN8ey6iaAA3kQ**"
	EmailAuthorizationCode = "bzwosnofbjsgiijb"
	EmailServerAddr        = "smtp.qq.com:25"
	Sender                 = "1519695805@qq.com"
	EmailHost              = "smtp.qq.com"
	EmailSubject           = "验证码"
	DefaultAvatar          = "https://www.static.talkxj.com/avatar/user.png"
	IdentityKey            = "blog-gin"
)

//RedisPrefixConst
const (
	CodeExpireTime    = 15 * 60 * 1000         //验证码过期时间
	CodeKey           = "code_"                //验证码
	BlogViewsCount    = "blog_views_count"     //博客浏览量
	ArticleViewsCount = "article_views_count_" //文章浏览量
	ArticleLikeCount  = "article_like_count"   //文章点赞量
	ArticleUserLike   = "article_user_like"    //用户点赞文章
	CommentLikeCount  = "comment_like_count"   //评论点赞量
	CommentUserLike   = "comment_user_like"    //用户点赞评论
	ABOUT             = "about"                //关于我信息
	NOTICE            = "notice"               //公告
	IpSet             = "ip_set"               //ip集合
	ChatServer        = "chat_server"          //聊天服
	Component         = "Layout"               //前端组件名
)

// SuccessOK code status
const (
	SuccessOK = 1000
)

//urls
const (
	BlogBaseUrl         = "/blog"
	BlogInfoUrl         = "/blog_info"
	ArticleList         = "/articles"
	ArticleById         = "/articles/:id"
	Tags                = "/tags"
	Archive             = "/archives"
	Categories          = "/categories"
	ArticleByCategoryID = "/categories/:categoryId"
	ArticleByTagID      = "/tags/:tagId"
	About               = "/about"
	Message             = "/messages"
	FriendLinks         = "/links"
	Comments            = "/comments"
	Replies             = "/replies/:commentId"
	LikeArticle         = "/articles/like"

	UserBaseUrl = "/users"
	VerifyCode  = "/code"
	Register    = "/register"
	Login       = "/login"

	AdminBaseUrl    = "/admin"
	UserMenu        = "/user/menus"
	AdminHomeData   = ""
	ArticleOptions  = "/articles/options"
	UploadImage     = "/articles/images"
	Articles        = "/articles"
	RefreshToken    = "/refresh_token"
	ArticlesByID    = "/articles/:id"
	ArticlesStatus  = "/articles/status"
	ArticleTop      = "/articles/top/:id"
	AdminCategories = "/categories"
	AdminTags       = "/tags"
)
