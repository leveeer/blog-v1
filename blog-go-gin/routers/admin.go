package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func adminRouters(r *gin.Engine) {
	jwtMiddleware := middleware.NewJWT()
	authMiddleware := jwtMiddleware.GinJWTMiddlewareInit()
	admin := r.Group(common.AdminBaseUrl)
	admin.POST(common.Login, authMiddleware.LoginHandler)
	//admin.Use(authMiddleware.MiddlewareFunc())
	{

	}
}
