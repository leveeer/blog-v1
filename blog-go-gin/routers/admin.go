package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/adminApi"
	"blog-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

var (
	MenuApi = adminApi.NewMenuRestApi()
)

func adminRouters(r *gin.Engine) {
	jwtMiddleware := middleware.NewJWT()
	authMiddleware := jwtMiddleware.GinJWTMiddlewareInit()
	admin := r.Group(common.AdminBaseUrl)
	admin.POST(common.Login, authMiddleware.LoginHandler)
	admin.GET(common.RefreshToken, authMiddleware.RefreshHandler) //刷新token
	admin.Use(authMiddleware.MiddlewareFunc(), middleware.AuthCheckRole())
	{
		admin.GET(common.UserMenu, MenuApi.GetUserMenus)
		admin.GET(common.AdminHomeData, BlogInfoApi.GetAdminHomeData)
		admin.GET(common.ArticleOptions, ArticleApi.GetArticleOptions)
		admin.POST(common.UploadImage, ArticleApi.UploadImage)
		admin.POST(common.Articles, ArticleApi.AddArticle)
		admin.GET(common.Articles, ArticleApi.GetAdminArticles)
		admin.PUT(common.Articles, ArticleApi.UpdateArticle)
		admin.DELETE(common.Articles, ArticleApi.DeleteArticles)
		admin.GET(common.ArticlesByID, ArticleApi.GetArticleInfoById)
		admin.PUT(common.ArticlesStatus, ArticleApi.UpdateArticleStatus)
		admin.PUT(common.ArticleTop, ArticleApi.UpdateArticleTop)
		admin.GET(common.AdminCategories, CategoryApi.GetAdminCategories)
		admin.POST(common.AdminCategories, CategoryApi.AddOrUpdateCategory)
		admin.PUT(common.AdminCategories, CategoryApi.AddOrUpdateCategory)
		admin.DELETE(common.AdminCategories, CategoryApi.DeleteCategory)
		admin.GET(common.AdminTags, TagApi.GetAdminTags)
		admin.POST(common.AdminTags, TagApi.AddOrUpdateTag)
		admin.PUT(common.AdminTags, TagApi.AddOrUpdateTag)
		admin.DELETE(common.AdminTags, TagApi.DeleteTag)
		admin.GET(common.AdminComments, CommentApi.GetAdminComments)
		admin.PUT(common.AdminComments, CommentApi.UpdateCommentStatus)
		admin.DELETE(common.AdminComments, CommentApi.DeleteComment)

		admin.GET(common.AdminMessages, MessageApi.GetAdminMessages)
		admin.DELETE(common.AdminMessages, MessageApi.DeleteMessage)

		admin.GET(common.AdminUsers, UserAuthApi.GetAdminUsers)
	}
}
