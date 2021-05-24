package web

import (
	"blog-go-gin/common"
	"blog-go-gin/controllers/base"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleRestApi struct {
	base.Controller
}

func (c *ArticleRestApi) GetArticleList(ctx *gin.Context) {
	var ipage page.IPage
	ipage.Current, _ = strconv.Atoi(ctx.Query("current"))
	articles, err := service.ArticleService.GetArticleList(ipage)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetArticlesFail)
		return
	}
	c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, &articles)
}

func (c *ArticleRestApi) GetArticleById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	article, err := service.ArticleService.GetArticleById(id)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetArticleByIdFail)
		return
	}
	c.RespSuccess(ctx,http.StatusOK, common.SuccessOK, &article)
}
