package blogApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	CategoryService = &service.CategoryService{}
)

type CategoryRestApi struct {
	base.Handler
}

func (c *CategoryRestApi) GetCategories(ctx *gin.Context) {
	categories, err := CategoryService.GetCategories()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetCategoriesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Categories: categories,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}

func (c *CategoryRestApi) GetArticleByCategoryID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("categoryId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	articles, err := ArticleService.GetArticleByCategoryID(id, &page.IPage{Current: currentPage})
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetArticleByCategoryIDFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:                   pb.Response_ResponseBeginIndex,
		Code:                    pb.ResultCode_SuccessOK,
		ServerTime:              time.Now().Unix(),
		ArticlesByCategoryOrTag: articles,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}
