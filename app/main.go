package main

import (
	"context"
	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/server"
	"github.com/yezihack/gofun/app/tools"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	k3log.NewDevelopment(server.Config.Title, tools.GetCurrentDirectory()+"/gofun.log")
	k3log.Info(server.Config.Title + "运行中...")
	c := cron.NewWithLocation(config.BeijingLocation)
	server.Start(c)

	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGINT)
	<-sigChan
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	cancel()
	c.Stop()
}
