package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/server"
	"github.com/yezihack/gofun/app/tools"
	"syscall"
)

func main() {
	//设置日志
	k3log.NewDevelopment(server.Config.Title, tools.GetCurrentDirectory()+"/gofun.log")
	k3log.Info(server.Config.Title + "运行中...")
	c := cron.NewWithLocation(config.BeijingLocation)

	stopChan := make(chan struct{})
	//调用主程序
	server.Start(c)
	server.Watcher(stopChan)

	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGINT)
	<-sigChan

	stopChan <- struct{}{}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	cancel()
	c.Stop()
}
