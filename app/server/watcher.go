package server

import (
	"fmt"
	"github.com/ThreeKing2018/k3log"
	"github.com/fsnotify/fsnotify"
)

func Watcher(stop chan struct{}) (err error) {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		k3log.Error("Watcher", err)
		return
	}
	//defer watch.Close()
	//监控文件
	err = watch.Add(Serve.ConfigPath)
	if err != nil {
		k3log.Error(err)
	}
	k3log.Info("对文件进行监控...", Serve.ConfigPath)

	go func() {
		for {
			select {
			case event := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if event.Op&fsnotify.Remove == fsnotify.Remove ||
						event.Op&fsnotify.Rename == fsnotify.Rename ||
						event.Op&fsnotify.Write == fsnotify.Write ||
						event.Op&fsnotify.Create == fsnotify.Create {
						watch.Remove(event.Name)
						watch.Add(event.Name)
						fmt.Println(event.Name, event.Op.String())
						if !Serve.LoadConfig() {
							k3log.Warn("加载配置文件异常")
							return
						}
						Meal.ReInit()
					}
				}
			case err := <-watch.Errors:
				k3log.Error("watch", err)
				return
			}
		}
	}()
	//开个goroutine进行关闭监控句柄
	go func() {
		<-stop
		watch.Close()
		fmt.Println("Watcher Close")
	}()
	return
}
