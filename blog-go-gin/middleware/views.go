package middleware

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"github.com/gin-gonic/gin"
)

func StatisticalViews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取用户ip
		clientIP := ctx.ClientIP()
		isMember, err := common.GetRedisUtil().SIsMember(common.IpSet, clientIP)
		if err != nil {
			logging.Logger.Debug(err)
		}
		if !isMember {
			//将redis中的访问量+1
			common.GetRedisUtil().IncrBy(common.BlogViewsCount, 1)
			common.GetRedisUtil().SAdd(common.IpSet, clientIP)
		}
		ctx.Next()
	}
}
