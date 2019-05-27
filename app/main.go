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
	k3log.NewDevelopment("钉钉消息", "logs")
	k3log.Info("go fun运行中...")
	c := cron.New()

	//随机吃饭
	c.AddFunc("0 30 12 * * 1-5", func() {
		meal := server.Meal.Random()
		err := tools.DingDingCurlPost(config.DingDingUrl["meal"], tools.SetDingDingData(meal, true))
		if err != nil {
			k3log.Error(err)
		}
	})
	//一周回顾
	c.AddFunc("0 0 12 * * 5", func() {
		historyWeek := server.Meal.History() //历史回顾
		tools.DingDingCurlPost(config.DingDingUrl["meal"], tools.SetDingDingData(historyWeek, true))
	})
	//下班
	c.AddFunc("0 30 18 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.Off(), true))
	})
	c.AddFunc("0 45 18 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.Off(), true))
	})
	c.AddFunc("0 0 19 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.Off(), true))
	})
	//上班
	c.AddFunc("0 0 9 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.On(), true))
	})
	c.AddFunc("0 15 9 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.On(), true))
	})
	c.AddFunc("0 25 9 * * *", func() {
		office := server.Office{}
		tools.DingDingCurlPost(config.DingDingUrl["office"], tools.SetDingDingData(office.On(), true))
	})

	c.Start()

	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGINT)
	<-sigChan
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	cancel()
	c.Stop()
}
