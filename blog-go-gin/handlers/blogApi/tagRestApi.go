package blogApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TagRestApi struct {
	base.Handler
}

var (
	TagService service.ITagService = &impl.TagServiceImpl{}
)

func (c *TagRestApi) GetTags(ctx *gin.Context) {
	tags, err := TagService.GetTags()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetTagsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Tags:       tags,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}

func (c *TagRestApi) GetArticlesByTagID(ctx *gin.Context) {
	tagId, _ := strconv.Atoi(ctx.Param("tagId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	articles, err := ArticleService.GetArticleByTagID(tagId, &page.IPage{Current: currentPage})
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
