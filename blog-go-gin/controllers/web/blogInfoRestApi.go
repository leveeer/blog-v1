package web

import (
	"blog-go-gin/common"
	"blog-go-gin/controllers/base"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var BlogInfoRestApi = &blogInfoRestApi{}

type blogInfoRestApi struct {
	base.Controller
}

func (c *blogInfoRestApi) GetBlogInfo(ctx *gin.Context) {
	blogInfo := service.BlogInfoService.GetBlogInfo()
	c.Result(ctx, http.StatusOK, common.SuccessOK, blogInfo, "查询成功")
}
