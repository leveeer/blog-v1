package routers

import (
	conf "blog-go-gin/config"
	"blog-go-gin/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitWebRouter() *gin.Engine {
	config := conf.GetConf()
	r := gin.New()
	if config.RunMode == "dev" {
		gin.SetMode(gin.DebugMode)
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowMethods = []string{"GET", "POST", "HEAD", "OPTIONS", "PATCH", "DELETE", "PUT"}
		corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
		corsConfig.MaxAge = time.Hour
		r.Use(cors.New(corsConfig))
	} else {
		r.Use(func(c *gin.Context) {
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusOK)
			}
		})
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(middleware.LoggerMiddleware(), middleware.AppRecoveryMiddleware(), middleware.StatisticWebSiteViews())
	blogRouters(r)
	userRouters(r)
	adminRouters(r)
	return r
}
