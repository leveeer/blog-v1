package main

import (
	"blog-go-gin/chat/logic"
	"blog-go-gin/chat/router"
	"blog-go-gin/config"
	"blog-go-gin/logging"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	//"net/http"
	"blog-go-gin/common"
)

var cleanup *bool

func GetServerInfo() *config.WS {
	conf := config.GetConf().Ws
	logging.Logger.Printf("%s:%d", conf.Host, conf.Port)
	return &conf
}

func init() {
	//todo 动态选择
	cleanup = flag.Bool("cleanup", false, "cleanup flag")
	flag.Parse()
}

func main() {
	var mem runtime.MemStats

	ws := GetServerInfo()
	if *cleanup {
		router.UnregisterGameServer(ws)
		return
	}

	srv := &http.Server{Addr: fmt.Sprint(ws.Host+":", ws.Port)}
	http.HandleFunc("/", router.HandleWs)
	router.RegisterGameServer(ws)

	logging.Logger.Infoln("ListenAndServe: ", ws)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logging.Logger.Fatalf("ListenAndServe(): %s", err)
		}
	}()
	logic.InitChatLogic()

	router.UnregisterGameServer(ws)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ctx.Done()
	stop()
	logging.Logger.Infof("Shutdown Server gracefully by %s...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Fatal("Server Shutdown:", err) // failure/timeout shutting down the server gracefully
	}

	// TODO: use sync wait
	time.Sleep(1 * time.Second)
	close(router.ClosedChan)
	common.GracefulWorkerWait()
	// close GRpc service
	//s.GracefulStop()
	runtime.ReadMemStats(&mem)
	logging.Logger.Println(mem.Alloc)
	logging.Logger.Println(mem.TotalAlloc)
	logging.Logger.Println(mem.HeapAlloc)
	logging.Logger.Println(mem.HeapSys)
	logging.Logger.Infoln("ChatService gracefully stopped.")
}
