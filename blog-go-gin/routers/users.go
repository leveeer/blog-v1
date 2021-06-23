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
	blog := r.Group(common.UserBaseUrl)
	blog.GET(common.VerifyCode, UsersApi.GetLoginCode)
}
