package usersApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/logging"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	UserAuthService service.IUserAuthService = &impl.UserAuthServiceImpl{}
)

type UserAuthRestApi struct {
	base.Handler
}

func (c *UserAuthRestApi) Register(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	err = UserAuthService.Register(request.User)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.RegisterFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "注册成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *UserAuthRestApi) Login(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	err = UserAuthService.Login(request.User)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.RegisterFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "注册成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *UserAuthRestApi) GetLoginCode(ctx *gin.Context) {
	username := ctx.Query("username")
	err := UserAuthService.GetLoginCode(username)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetLoginCodeFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
