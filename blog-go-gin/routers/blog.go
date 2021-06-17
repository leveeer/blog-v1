package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/blogApi"
	"github.com/gin-gonic/gin"
)

var (
	ArticleApi      = &blogApi.ArticleRestApi{}
	BlogInfoApi     = &blogApi.BlogInfoRestApi{}
	TagApi          = &blogApi.TagRestApi{}
	CategoryRestApi = &blogApi.CategoryRestApi{}
)

func blogRouters(r *gin.Engine) {
	blog := r.Group(common.BlogBaseUrl)
	blog.GET(common.BlogInfoUrl, BlogInfoApi.GetBlogInfo)
	blog.POST(common.ArticleList, ArticleApi.GetArticleList)
	blog.GET(common.ArticleById, ArticleApi.GetArticleById)
	blog.POST(common.Archive, ArticleApi.GetArticleArchives)
	blog.POST(common.TagList, TagApi.GetTagList)
	blog.GET(common.Categories, CategoryRestApi.GetCategories)
	blog.GET(common.ArticleByCategoryID, CategoryRestApi.GetArticleByCategoryID)
}
