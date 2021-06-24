package base

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"github.com/gin-gonic/gin"
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
}

type Handler struct {
	Wg sync.WaitGroup
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
