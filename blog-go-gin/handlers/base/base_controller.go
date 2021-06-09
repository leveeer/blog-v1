package base

import (
	"blog-go-gin/common"
	"github.com/gin-gonic/gin"
	"sync"
)

// Base /*
type Base interface {
	RespSuccess(ctx *gin.Context, httpCode, code int, data interface{})
	RespFailWithDesc(ctx *gin.Context, httpCode int, code common.ErrorCode)
	ThrowError(code string, message string)
}

type Controller struct {
	Wg sync.WaitGroup
}

func (c *Controller) RespSuccess(ctx *gin.Context, httpCode, code int, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"success": true,
		"code":    code,
		"data":    data,
	})
}

func (c *Controller) RespFailWithDesc(ctx *gin.Context, httpCode int, code common.ErrorCode) {
	ctx.JSON(httpCode, gin.H{
		"success": false,
		"code":    code,
		"message": common.GetMsg(code),
	})
}

func (c *Controller) ThrowError(code string, message string) {

}
