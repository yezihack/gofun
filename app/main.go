package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/server"
	"github.com/yezihack/gofun/app/tools"
)

func main() {
	//设置日志
	k3log.NewDevelopment(server.Config.Title, tools.GetCurrentDirectory()+"/gofun.log")
	k3log.Info(server.Config.Title + "运行中...")
	c := cron.NewWithLocation(config.BeijingLocation)
	//调用主程序
	server.Start(c)
	server.Watcher()

	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGINT)
	<-sigChan
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	cancel()
	c.Stop()
}
