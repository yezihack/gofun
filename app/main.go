package main

import (
	"context"
	"time"

	"fmt"
	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/cron"
	"github.com/yezihack/gofun/app/config"
	"github.com/yezihack/gofun/app/server"
	"github.com/yezihack/gofun/app/tools"
	"net/http"
)

func main() {
	//设置日志
	k3log.NewDevelopment(server.Serve.Config.Title, tools.GetCurrentDirectory()+"/gofun.log")
	k3log.Info(server.Serve.Config.Title + "运行中...")
	c := cron.NewWithLocation(config.BeijingLocation)

	stopChan := make(chan struct{})
	//调用主程序
	server.Start(c, stopChan)
	server.Watcher(stopChan)

	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, os.Kill, os.Interrupt)
	//<-sigChan
	k3log.Info("启动成功, port", config.Port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", config.Port), nil)
	if err != nil {
		k3log.Error("端口占用", config.Port)
	}
	k3log.Warn("准备关闭服务")
	stopChan <- struct{}{}
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	cancel()
	c.Stop()
}
