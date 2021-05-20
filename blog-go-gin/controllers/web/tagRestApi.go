package web

import (
	"blog-go-gin/controllers/base"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TagRestApi struct {
	base.Controller
}

func (r *TagRestApi) GetTagList(ctx *gin.Context) {
	var pageObj page.IPage
	err := ctx.ShouldBindQuery(&pageObj)
	if err != nil {
		log.Println("BindQuery failed, err:", err)
	}

	tagList := service.GetTagList(pageObj)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"list":  tagList,
		"msg":   "查询成功",
		"count": len(tagList),
	})
}
