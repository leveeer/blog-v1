package web

import (
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleRestApi struct {
}

func (c *ArticleRestApi) GetArticleList(ctx *gin.Context) {
	var ipage page.IPage

	err := ctx.ShouldBindJSON(&ipage)
	if err != nil {
		fmt.Println("GetArticleList ShouldBindJSON err:", err)
	}

	fmt.Println(ipage)
	articles := service.ArticleService.GetArticleList(ipage)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"data":    articles,
		"total":   len(articles),
	})
}

func (c *ArticleRestApi) GetArticleByUid(ctx *gin.Context) {

}
