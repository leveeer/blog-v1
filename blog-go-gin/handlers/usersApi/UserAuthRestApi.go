package usersApi

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
	UserAuthService service.IUserAuthService = &impl.UserAuthServiceImpl{}
)

type UserAuthRestApi struct {
	base.Handler
}

func (c *UserAuthRestApi) GetLoginCode(ctx *gin.Context) {
	username := ctx.Query("username")
	err := UserAuthService.GetLoginCode(username)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetLoginCodeFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}
