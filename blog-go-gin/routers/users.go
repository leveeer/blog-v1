package routers

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/usersApi"
	"github.com/gin-gonic/gin"
)

var (
	UsersApi = &usersApi.UserAuthRestApi{}
)

func userRouters(r *gin.Engine) {
	user := r.Group(common.UserBaseUrl)
	user.GET(common.VerifyCode, UsersApi.GetLoginCode)
	user.POST(common.Register, UsersApi.Register)
	user.POST(common.Login, UsersApi.Login)
}
