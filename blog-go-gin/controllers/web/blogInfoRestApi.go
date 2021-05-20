package web

import (
	"blog-go-gin/common"
	"blog-go-gin/controllers/base"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogInfoRestApi struct {
	base.Controller
}

func (c *BlogInfoRestApi) GetBlogInfo(ctx *gin.Context) {
	blogHomeInfoVo, err := service.BlogInfoService.GetBlogInfo()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetBlogHomeInfoFail)
	}
	c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, blogHomeInfoVo)
}
