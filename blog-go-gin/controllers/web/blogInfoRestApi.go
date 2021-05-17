package web

import (
	"blog-go-gin/controllers/base"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogInfoRestApi struct {
	base.Controller
}

func (c *BlogInfoRestApi) GetBlogInfo(ctx *gin.Context) {
	blogInfo := service.BlogInfoService.GetBlogInfo()
	ctx.JSON(http.StatusOK, base.BaseController.Result(http.StatusOK, blogInfo, "查询成功"))
}
