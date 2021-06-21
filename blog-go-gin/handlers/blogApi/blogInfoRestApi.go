package blogApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	BlogInfoService service.IBlogInfoService = &impl.BlogInfoServiceImpl{}
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
		Code:         pb.ResultCode_SuccessOK,
		ServerTime:   time.Now().Unix(),
		BlogHomeInfo: blogHomeInfo,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}

func (c *BlogInfoRestApi) GetAbout(ctx *gin.Context) {

	about, err := BlogInfoService.GetAbout()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetAboutFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		About:      about,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)

}
