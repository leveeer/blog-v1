package main

import (
	"blog-go-gin/common"
	"blog-go-gin/config"
	"blog-go-gin/dao"
	"blog-go-gin/jobs"
	"blog-go-gin/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	gin.ForceConsoleColor()                            // 启用gin的日志输出带颜色
	gin.DefaultWriter = colorable.NewColorableStdout() // 替换默认Writer（关键步骤）
	var router *gin.Engine
	dao.InitMysql()
	common.InitRedis()
	router = routers.InitWebRouter()
	//注册定时任务
	jobs.RegisterCron()
	_ = router.Run(fmt.Sprintf(":%d", config.GetConf().HttpPort))
}
