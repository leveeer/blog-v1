package main

import (
	"blog-go-gin/chat/router"
	"blog-go-gin/config"
	"blog-go-gin/logging"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ws := GetServerInfo()
	if *cleanup {
		router.UnregisterGameServer(ws)
		return
	}

	//s := initGRpcService(config)

	srv := &http.Server{Addr: fmt.Sprint(ws.Host+":", ws.Port)}
	http.HandleFunc("/", router.HandleWs)
	router.RegisterGameServer(ws)

	logging.Logger.Infoln("ListenAndServe: ", ws)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// TODO: there is a chance that next line won't have time to run,
			// as main() doesn't wait for this goroutine to stop. don't use
			// code with race conditions like these for production. see post
			// comments below on more discussion on how to handle this.
			logging.Logger.Fatalf("ListenAndServe(): %s", err)
		}
		logging.Logger.Infoln("http server shutdown...")
	}()

	sig := <-stop
	logging.Logger.Infoln("Shutting down the server... by ", sig)
	router.UnregisterGameServer(ws)

	// TODO: now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	// TODO: use sync wait
	time.Sleep(1 * time.Second)
	close(make(chan struct{}))
	common.GracefulWorkerWait()
	// close GRpc service
	//s.GracefulStop()
	runtime.ReadMemStats(&mem)
	logging.Logger.Println(mem.Alloc)
	logging.Logger.Println(mem.TotalAlloc)
	logging.Logger.Println(mem.HeapAlloc)
	logging.Logger.Println(mem.HeapSys)
	logging.Logger.Infoln("GameService gracefully stopped.")

}
