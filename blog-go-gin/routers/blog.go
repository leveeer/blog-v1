package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/blogApi"
	"github.com/gin-gonic/gin"
)

var (
	ArticleApi    = &blogApi.ArticleRestApi{}
	BlogInfoApi   = &blogApi.BlogInfoRestApi{}
	TagApi        = &blogApi.TagRestApi{}
	CategoryApi   = &blogApi.CategoryRestApi{}
	MessageApi    = &blogApi.MessageRestApi{}
	FriendLinkApi = &blogApi.FriendLinkRestApi{}
	CommentApi    = &blogApi.CommentRestApi{}
)

func blogRouters(r *gin.Engine) {
	blog := r.Group(common.BlogBaseUrl)
	blog.GET(common.BlogInfoUrl, BlogInfoApi.GetBlogInfo)
	blog.POST(common.ArticleList, ArticleApi.GetArticleList)
	blog.GET(common.ArticleById, ArticleApi.GetArticleById)
	blog.POST(common.Archive, ArticleApi.GetArticleArchives)
	blog.GET(common.Tags, TagApi.GetTags)
	blog.GET(common.Categories, CategoryApi.GetCategories)
	blog.GET(common.ArticleByCategoryID, CategoryApi.GetArticleByCategoryID)
	blog.GET(common.ArticleByTagID, TagApi.GetArticlesByTagID)
	blog.GET(common.About, BlogInfoApi.GetAbout)
	blog.GET(common.Message, MessageApi.GetMessages)
	blog.POST(common.Message, MessageApi.AddMessages)
	blog.GET(common.FriendLinks, FriendLinkApi.GetFriendLinks)
	blog.GET(common.Comments, CommentApi.GetComments)
}
