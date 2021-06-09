package web

import (
	"blog-go-gin/common"
	"blog-go-gin/handlers/base"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	BlogInfoService = &service.BlogInfoService{}
)

type BlogInfoRestApi struct {
	base.Controller
}

func (c *BlogInfoRestApi) GetBlogInfo(ctx *gin.Context) {
	blogHomeInfoVo, err := BlogInfoService.GetBlogInfo()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetBlogHomeInfoFail)
		return
	}
	c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, blogHomeInfoVo)
}
