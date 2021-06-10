package web

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	BlogInfoService = &service.BlogInfoService{}
)

type BlogInfoRestApi struct {
	base.Handler
}

func (c *BlogInfoRestApi) GetBlogInfo(ctx *gin.Context) {
	blogHomeInfo, err := BlogInfoService.GetBlogInfo()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetBlogHomeInfoFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:        pb.Response_ResponseBeginIndex,
		Code:         pb.ResultCode_Success,
		ServerTime:   time.Time.Unix(time.Now()),
		BlogHomeInfo: blogHomeInfo,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}
