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

var (
	CommentService service.ICommentService = &impl.CommentServiceImpl{}
)

type CommentRestApi struct {
	base.Handler
}

func (c *CommentRestApi) GetComments(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("articleId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	commentInfo, err := CommentService.GetComments(id, &page.IPage{Current: currentPage})
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetCommentsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		CommentInfo: commentInfo,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}
