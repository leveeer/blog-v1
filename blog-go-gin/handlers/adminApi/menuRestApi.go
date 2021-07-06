package adminApi

import (
	"blog-go-gin/handlers/base"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
)

var (
	MenuService service.IMenuService = impl.NewMenuServiceImpl()
)

type MenuRestApi struct {
	base.Handler
}

func NewMenuRestApi() *MenuRestApi {
	return &MenuRestApi{}
}

func (c *MenuRestApi) GetUserMenu(ctx *gin.Context) {

}
