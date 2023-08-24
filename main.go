package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baaj2109/shorturl/common"
	"github.com/baaj2109/shorturl/router"
)

func init() {
	common.InitLogger()
	common.InitConfig()
	common.InitDB()
}

func main() {
	engine := router.InitRouter()

	server := http.Server{
		Addr:    common.Config.GetString("app.host") + ":" + common.Config.GetString("app.port"),
		Handler: engine,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.SugarLogger.Fatal("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit,
		syscall.SIGTSTP, syscall.SIGUSR1, syscall.SIGUSR2,
		syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		common.SugarLogger.Fatal("server shutdown err:%s", err)
	}
	common.SugarLogger.Info("server shutdown success")
}
