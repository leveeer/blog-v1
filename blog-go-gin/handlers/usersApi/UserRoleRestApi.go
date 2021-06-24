package usersApi

import (
	"blog-go-gin/handlers/base"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
)

var (
	UserRoleService service.IUserRoleService = &impl.UserRoleServiceImpl{}
)

type UserRoleRestApi struct {
	base.Handler
}
