package middleware

import (
	"blog-go-gin/common"
	"github.com/gin-contrib/sessions" // 导入session包
	"github.com/gin-gonic/gin"
)

func StatisticalViews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取用户ip
		clientIP := common.GetClientIP(ctx)
		//从session中获取ip
		// 初始化session对象
		session := sessions.Default(ctx)
		ip := session.Get("ip")
		if ip == nil || clientIP != ip.(string) {
			//没有，说明用户未访问过
			//将ip存入session
			session.Set("ip", clientIP)
			//一定要save!!!
			_ = session.Save()
			//将redis中的访问量+1
			common.RedisUtil.IncrBy(common.BlogViewsCount, 1)
		}
		//有,说明用户已访问
		common.RedisUtil.SAdd(common.IpSet, clientIP)
		ctx.Next()
	}
}
