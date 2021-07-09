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
	admin.GET("/refresh_token", authMiddleware.RefreshHandler) //刷新token
	admin.Use(authMiddleware.MiddlewareFunc(), middleware.AuthCheckRole())
	{
		admin.GET(common.UserMenu, MenuApi.GetUserMenus)
		admin.GET(common.AdminHomeData, BlogInfoApi.GetAdminHomeData)
		admin.GET(common.ArticleOptions, ArticleApi.GetArticleOptions)
		admin.POST(common.UploadImage, ArticleApi.UploadImage)
		admin.POST(common.Articles, ArticleApi.AddArticle)
	}
}