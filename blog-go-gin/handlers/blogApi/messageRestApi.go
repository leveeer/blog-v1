package blogApi

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

type MessageRestApi struct {
	base.Handler
}

var (
	MessageService service.IMessageService = &impl.MessageServiceImpl{}
)

func (c *MessageRestApi) GetMessages(ctx *gin.Context) {
	messages, err := MessageService.GetMessages()
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetMessagesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Messages:   messages,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}

func (c *MessageRestApi) AddMessages(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.InvalidRequestParams)
		logging.Logger.Error(err)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.RespFailWithDesc(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	request.Message.IpAddress = ctx.ClientIP()
	request.Message.CreateTime = time.Now().Unix()
	err = MessageService.AddMessage(request.Message)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.AddMessageFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
}
