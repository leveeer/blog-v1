package middleware

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StatisticWebSiteViews() gin.HandlerFunc {
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

func StatisticArticleViews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取用户ip
		clientIP := ctx.ClientIP()
		articleID := ctx.Param("id")
		err := common.GetRedisUtil().PFAdd(common.ArticleViewsCount+articleID, clientIP)
		if err != nil {
			logging.Logger.Debug(err)
		}
		id, _ := strconv.Atoi(articleID)
		err = model.UpdateArticleClickCount(id)
		if err != nil {
			logging.Logger.Debug(err)
		}
		ctx.Next()
	}
}
