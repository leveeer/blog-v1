package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/controllers/web"
	"github.com/gin-gonic/gin"
)

var (
	ArticleApi  = &web.ArticleRestApi{}
	BlogInfoApi = &web.BlogInfoRestApi{}
	TagApi      = &web.TagRestApi{}
)

func blogRouters(r *gin.Engine) {
	blog := r.Group(common.BlogBaseUrl)
	blog.GET(common.BlogInfoUrl, BlogInfoApi.GetBlogInfo)
	blog.GET(common.ArticleList, ArticleApi.GetArticleList)
	blog.GET(common.ArticleById, ArticleApi.GetArticleById)
	blog.GET(common.TagList, TagApi.GetTagList)
}
