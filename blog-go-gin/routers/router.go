package routers

import (
	conf "blog-go-gin/config"
	"blog-go-gin/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	// 创建基于cookie的存储引擎，secret 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	// 设置session中间件，参数session，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，可以替换成其他存储引擎
	r.Use(middleware.LoggerMiddleware(), middleware.AppRecoveryMiddleware(), sessions.Sessions("session", store), middleware.StatisticalViews())
	blogRouters(r)
	userRouters(r)
	adminRouters(r)
	return r
}
