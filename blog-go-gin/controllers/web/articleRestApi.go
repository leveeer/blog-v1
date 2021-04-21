package web

import (
	"blog-go-gin/models/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleRestApi struct {

}

/*
	/getArticleList?keyword=&likes=&state=1&tag_id=&category_id=&pageNum=1&pageSize=10
*/
func (c *ArticleRestApi) GetArticleList(context *gin.Context) {
	var articleVO vo.ArticleVO
	err := context.ShouldBindQuery(&articleVO)
	if err != nil {
		fmt.Println("BindQuery failed, err:", err)
	}
	fmt.Println(&articleVO)
	context.JSON(http.StatusOK, gin.H{
		"data": articleVO,
	})
}
