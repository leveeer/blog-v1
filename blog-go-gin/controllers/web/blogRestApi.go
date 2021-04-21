package web

import (
	"blog-go-gin/models/vo"
	"blog-go-gin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogRestApi struct {
}

/*
	/getArticleList?keyword=&likes=&state=1&tag_id=&category_id=&pageNum=1&pageSize=10
*/
func (c *BlogRestApi) GetBlogList(context *gin.Context) {
	var blogVO vo.BlogVO
	err := context.ShouldBindQuery(&blogVO)
	if err != nil {
		fmt.Println("BindQuery failed, err:", err)
	}

	fmt.Println(blogVO)

	blogList := service.GetBlogList(blogVO)
	context.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"list":  blogList,
		"msg":   "查询成功",
		"count": len(blogList),
	})
}
