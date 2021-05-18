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
var ArticleRestApi = &articleRestApi{}

type articleRestApi struct {
	base.Controller
}

func (c *articleRestApi) GetArticleList(ctx *gin.Context) {
	var ipage page.IPage
	ipage.Current, _ = strconv.Atoi(ctx.Query("current"))
	articles := service.ArticleService.GetArticleList(ipage)
	c.Result(ctx, http.StatusOK, common.SuccessOK, articles, "查询成功")
}

func (c *articleRestApi) GetArticleByUid(ctx *gin.Context) {

}
