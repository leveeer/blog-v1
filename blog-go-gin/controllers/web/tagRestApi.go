package web

import (
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagRestApi struct {

}

func (r *TagRestApi) GetTagList(ctx *gin.Context) {
	var pageObj page.IPage
	err := ctx.ShouldBindQuery(&pageObj)
	if err != nil {
		fmt.Println("BindQuery failed, err:", err)
	}

	tagList := service.GetTagList(pageObj)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"list":  tagList,
		"msg":   "查询成功",
		"count": len(tagList),
	})
}
