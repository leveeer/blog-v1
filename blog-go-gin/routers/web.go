package routers

import (
	"blog-go-gin/controllers/web"
	"blog-go-gin/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	ArticleRestApi *web.ArticleRestApi
)

func InitWebRouter(config *models.Config) *gin.Engine {
	r := gin.New()
	if config.RunMode == "dev" {
		gin.SetMode(gin.DebugMode)
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowMethods = []string{"GET", "POST", "HEAD", "OPTIONS", "PATCH", "DELETE", "PUT"}
		corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
		corsConfig.MaxAge = time.Hour
		r.Use(cors.New(corsConfig))
	}else {
		r.Use(func(c *gin.Context) {
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusOK)
			}
		})
		gin.SetMode(gin.ReleaseMode)
	}

	//   /getArticleList?keyword=&likes=&state=1&tag_id=&category_id=&pageNum=1&pageSize=10
	r.GET("/getArticleList", ArticleRestApi.GetArticleList)

	return r
}