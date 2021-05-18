package main

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	"blog-go-gin/jobs"
	"blog-go-gin/models"
	"blog-go-gin/routers"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"log"
	"strconv"
	"sync"
)

var OnceDo = sync.Once{}
var Conf *models.Config

func main() {
	gin.ForceConsoleColor()                            // 启用gin的日志输出带颜色
	gin.DefaultWriter = colorable.NewColorableStdout() // 替换默认Writer（关键步骤）
	var router *gin.Engine
	OnceDo.Do(func() {
		var config models.Config
		Conf = config.GetConf()
	})
	dao.InitMysql(Conf)
	err := common.InitRedis(Conf)
	if err != nil {
		log.Println("connect to redis failed, err:", err)
	}
	router = routers.InitWebRouter(Conf)
	//注册定时任务
	jobs.RegisterCron()
	_ = router.Run(":" + strconv.Itoa(int(Conf.HttpPort)))
}
