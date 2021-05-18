package base

import (
	"blog-go-gin/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

/*
基类
*/
type Controller struct {
	Wg sync.WaitGroup
}

func (c *Controller) Result(ctx *gin.Context, httpCode, code int, data interface{}, message string) {
	ctx.JSON(httpCode, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func RespSuccess(data interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func RespFailWithCode(code common.ErrorCode) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"msg":     common.GetMsg(code),
	}
}

func RespFailWithMsg(code int, msg string) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"msg":     msg,
	}
}

func RespFailWithDesc(code common.ErrorCode, msg string) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"msg":     fmt.Sprintf("%s[%s]", common.GetMsg(code), msg),
	}
}

func (c *Controller) ThrowError(code string, message string) {

}
