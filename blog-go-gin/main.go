package main

import (
	"blog-go-gin/models"
	"blog-go-gin/routers"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable" // 改包可完美解决问题
	"strconv"
	"sync"
)

var OnceDo = sync.Once{}
var Conf *models.Config


func main() {
	// 启用gin的日志输出带颜色
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()

	var router *gin.Engine
	OnceDo.Do(func() {
		var config models.Config
		Conf = config.GetConf()
		router = routers.InitWebRouter(Conf)
	})

	_ = router.Run(":" + strconv.Itoa(int(Conf.HttpPort)))

}

