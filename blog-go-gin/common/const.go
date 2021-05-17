package common

const (
	BloggerId = 1 //博主id
	False     = 0 //否
	True      = 1 //是
)

//RedisPrefixConst
const (
	CODE_EXPIRE_TIME    = 15 * 60 * 1000        //验证码过期时间
	CODE_KEY            = "code_"               //验证码
	BLOG_VIEWS_COUNT    = "blog_views_count"    //博客浏览量
	ARTICLE_VIEWS_COUNT = "article_views_count" //文章浏览量
	ARTICLE_LIKE_COUNT  = "article_like_count"  //文章点赞量
	ARTICLE_USER_LIKE   = "article_user_like"   //用户点赞文章
	COMMENT_LIKE_COUNT  = "comment_like_count"  //评论点赞量
	COMMENT_USER_LIKE   = "comment_user_like"   //用户点赞评论
	ABOUT               = "about"               //关于我信息
	NOTICE              = "notice"              //公告
	IP_SET              = "ip_set"              //ip集合

)
