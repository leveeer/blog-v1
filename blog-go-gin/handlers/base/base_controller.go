package base

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"sync"
	"time"
)

// Base /*
type Base interface {
	RespSuccess(ctx *gin.Context, httpCode, code int, data interface{})
	RespFailWithDesc(ctx *gin.Context, httpCode int, code common.ErrorCode)
	ThrowError(code string, message string)
	ProtoBuf(ctx *gin.Context, httpCode int, data interface{})
	ProtoBufFail(ctx *gin.Context, httpCode int, code common.ErrorCode)
	ReadRequestBody(ctx *gin.Context) (*pb.RequestPkg, error)
}

type Handler struct {
	Wg sync.WaitGroup
}

func (c *Handler) ReadRequestBody(ctx *gin.Context) (*pb.RequestPkg, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		return nil, err
	}
	logging.Logger.Debug(request)
	return request, nil
}

func (c *Handler) RespSuccess(ctx *gin.Context, httpCode, code int, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"success": true,
		"code":    code,
		"data":    data,
	})
}

func (c *Handler) RespFailWithDesc(ctx *gin.Context, httpCode int, code common.ErrorCode) {
	ctx.JSON(httpCode, gin.H{
		"success": false,
		"code":    code,
		"message": common.GetMsg(code),
	})
}

func (c *Handler) ThrowError(code string, message string) {

}

func (c *Handler) ProtoBuf(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.ProtoBuf(httpCode, data)
}

func (c *Handler) ProtoBufFail(ctx *gin.Context, httpCode int, code common.ErrorCode) {
	data := &pb.ResponsePkg{
		Code:       pb.ResultCode_Fail,
		ServerTime: time.Now().Unix(),
		Message:    common.GetMsg(code),
	}
	ctx.ProtoBuf(httpCode, data)
}
