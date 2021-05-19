package common

const (
	BloggerId = 1 //博主id
	False     = 0 //否
	True      = 1 //是
)

//RedisPrefixConst
const (
	CodeExpireTime    = 15 * 60 * 1000        //验证码过期时间
	CodeKey           = "code_"               //验证码
	BlogViewsCount    = "blog_views_count"    //博客浏览量
	ArticleViewsCount = "article_views_count" //文章浏览量
	ArticleLikeCount  = "article_like_count"  //文章点赞量
	ArticleUserLike   = "article_user_like"   //用户点赞文章
	CommentLikeCount  = "comment_like_count"  //评论点赞量
	CommentUserLike   = "comment_user_like"   //用户点赞评论
	ABOUT             = "about"               //关于我信息
	NOTICE            = "notice"              //公告
	IpSet             = "ip_set"              //ip集合
)

//code status
const (
	SuccessOK = 1000
)

//urls
const (
	BlogBaseUrl = "/blog"
	BlogInfoUrl = "/blogInfo"
	ArticleList = "/articles"
	ArticleById = "/articleById"
	TagList     = "/tags"
)
