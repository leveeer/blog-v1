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
	FriendLinkService service.IFriendLinkService = impl.NewFriendLinkServiceImpl()
)

type FriendLinkRestApi struct {
	base.Handler
}

func NewFriendLinkRestApi() *FriendLinkRestApi {
	return &FriendLinkRestApi{}
}

func (c *FriendLinkRestApi) GetFriendLinks(ctx *gin.Context) {
	friendLinks, err := FriendLinkService.GetFriendLinks()
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetMessagesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		FriendLinks: friendLinks,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)

}
