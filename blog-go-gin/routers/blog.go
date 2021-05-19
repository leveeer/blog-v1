package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/controllers/web"
	"github.com/gin-gonic/gin"
)

func blogRouters(r *gin.Engine) {
	blog := r.Group(common.BlogBaseUrl)
	blog.GET(common.BlogInfoUrl, web.BlogInfoRestApi.GetBlogInfo)
	blog.GET(common.ArticleList, web.ArticleRestApi.GetArticleList)
	blog.POST(common.ArticleById, web.ArticleRestApi.GetArticleById)
	blog.GET(common.TagList, web.TagRestApi.GetTagList)
}
