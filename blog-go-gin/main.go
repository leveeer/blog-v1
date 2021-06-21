package main

import (
	"blog-go-gin/common"
	"blog-go-gin/config"
	"blog-go-gin/dao"
	"blog-go-gin/jobs"
	"blog-go-gin/logging"
	"blog-go-gin/routers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(int(config.GetConf().HttpPort)),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ctx.Done()
	stop()
	logging.Logger.Infof("Shutdown Server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Fatal("Server Shutdown:", err)
	}
	logging.Logger.Infof("Server exiting")
}
