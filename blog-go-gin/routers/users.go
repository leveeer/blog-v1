package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/usersApi"
	"blog-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

var (
	UserAuthApi = usersApi.NewUserAuthRestApi()
)

func userRouters(r *gin.Engine) {
	jwtMiddleware := middleware.NewJWT()
	authMiddleware := jwtMiddleware.GinJWTMiddlewareInit()
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	user := r.Group(common.UserBaseUrl)
	{
		user.GET(common.VerifyCode, UserAuthApi.GetLoginCode)
		user.POST(common.Register, UserAuthApi.Register)
		user.POST(common.Login, authMiddleware.LoginHandler)
	}
}
