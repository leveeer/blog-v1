package routers

import (
	"blog-go-gin/controllers/web"
	"github.com/gin-gonic/gin"
)

func blogRouters(r *gin.Engine) {
	blog := r.Group("/blog")
	blog.GET("/blogInfo", web.BlogInfoRestApi.GetBlogInfo)
	blog.GET("/getArticleList", web.ArticleRestApi.GetArticleList)
	blog.POST("/getArticleDetail", web.ArticleRestApi.GetArticleByUid)
	blog.GET("/getTagList", web.TagRestApi.GetTagList)
}