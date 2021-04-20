package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleRestApi struct {
}

func (c *ArticleRestApi) GetArticleList(context *gin.Context) {
	fmt.Println("GetArticleList")
	context.JSON(http.StatusOK, gin.H{
		"message": "123123124234",
	})
}
