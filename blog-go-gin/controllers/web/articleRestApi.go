package web

import (
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
	articles := service.ArticleService.GetArticleList(ipage)
	ctx.JSON(http.StatusOK, base.BaseController.Result(http.StatusOK, articles, "查询成功"))
}

func (c *ArticleRestApi) GetArticleByUid(ctx *gin.Context) {

}
